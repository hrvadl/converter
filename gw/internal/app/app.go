package app

//nolint:revive
import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/hrvadl/converter/gw/docs"
	"github.com/hrvadl/converter/gw/internal/cfg"
	"github.com/hrvadl/converter/gw/internal/transport/grpc/clients/ratewatcher"
	ssvc "github.com/hrvadl/converter/gw/internal/transport/grpc/clients/sub"
	"github.com/hrvadl/converter/gw/internal/transport/http/handlers/rate"
	"github.com/hrvadl/converter/gw/internal/transport/http/handlers/sub"
	"github.com/hrvadl/converter/gw/pkg/logger"
)

const operation = "app init"

// New constructs new App with provided arguments.
// NOTE: than neither cfg or log can't be nil or App will panic.
func New(cfg cfg.Config, log *slog.Logger) *App {
	return &App{
		cfg: cfg,
		log: log,
	}
}

// App is a thin abstraction used to initialize all the dependencies,
// db connections, and GRPC server/clients. Could return an error if any
// of described above steps failed.
type App struct {
	cfg cfg.Config
	log *slog.Logger
}

// MustRun is a wrapper around App.Run() function which could be handly
// when it's called from the main goroutine and we don't need to handler
// an error.
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

// Run method creates new GRPC server then initializes MySQL DB connection,
// after that initializes all neccessary domain related services and finally
// starts listening on the provided ports. Could return an error if any of
// described above steps failed
func (a *App) Run() error {
	rw, err := ratewatcher.NewClient(
		a.cfg.RateWatcherAddr,
		a.log.With("source", "rateWatcherClient"),
	)
	if err != nil {
		return fmt.Errorf("%s: failed to initialize ratewatcher client: %w", operation, err)
	}

	subsvc, err := ssvc.NewClient(a.cfg.SubAddr, a.log.With("source", "subClient"))
	if err != nil {
		return fmt.Errorf("%s: failed to init sub service: %w", operation, err)
	}

	sh := sub.NewHandler(subsvc, a.log.With("source", "subHandler"))
	rh := rate.NewHandler(rw, a.log.With("source", "rateHandler"))

	r := chi.NewRouter()
	r.Use(
		middleware.Heartbeat("/health"),
		middleware.Recoverer,
		middleware.Logger,
		middleware.CleanPath,
		middleware.SetHeader("Content-Type", "application/octet-stream"),
	)

	r.Route("/api", func(r chi.Router) {
		r.Get("/rate", rh.GetRate)
		r.With(
			middleware.AllowContentType("application/x-www-form-urlencoded"),
		).Post("/subscribe", sh.Subscribe)
	})

	if a.cfg.LogLevel == "DEBUG" {
		r.Get("/docs/*", httpSwagger.WrapHandler)
	}

	a.log.Info("Starting web server", "addr", a.cfg.Addr)
	srv := newServer(
		r,
		a.cfg.Addr,
		slog.NewLogLogger(a.log.Handler(), logger.MapLevels(a.cfg.LogLevel)),
	)

	return srv.ListenAndServe()
}

// GracefulStop method gracefully stop the server. It listens to the OS sigals.
// After it recieves signal it terminates all currently active servers,
// client, connections (if any) and gracefully exits.
func (a *App) GracefulStop() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	signal := <-ch
	a.log.Info("Recieved stop signal. Terminating...", "signal", signal)
	a.log.Info("Successfully terminated server. Bye!")
}
