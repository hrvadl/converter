package app

import (
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"github.com/hrvadl/converter/sub/internal/cfg"
	"github.com/hrvadl/converter/sub/internal/service/cron"
	"github.com/hrvadl/converter/sub/internal/service/sender"
	"github.com/hrvadl/converter/sub/internal/service/sender/formatter"
	subs "github.com/hrvadl/converter/sub/internal/service/sub"
	"github.com/hrvadl/converter/sub/internal/service/validator"
	"github.com/hrvadl/converter/sub/internal/storage/platform/db"
	"github.com/hrvadl/converter/sub/internal/storage/subscriber"
	"github.com/hrvadl/converter/sub/internal/transport/grpc/clients/mailer"
	"github.com/hrvadl/converter/sub/internal/transport/grpc/clients/ratewatcher"
	"github.com/hrvadl/converter/sub/internal/transport/grpc/server/sub"
	"github.com/hrvadl/converter/sub/pkg/logger"
)

const operation = "app init"

const (
	cronJobHour   = 12
	cronJobMinute = 0
)

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
	srv *grpc.Server
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
	a.srv = grpc.NewServer(grpc.ChainUnaryInterceptor(
		logger.NewServerGRPCMiddleware(a.log),
	))

	db, err := db.NewConn(a.cfg.Dsn)
	if err != nil {
		return fmt.Errorf("%s: failed to init db: %w", operation, err)
	}

	sr := subscriber.NewRepo(db)
	v := validator.NewStdlib()
	svc := subs.NewService(sr, v)
	sub.Register(a.srv, svc, a.log.With("source", "sub"))

	m, err := mailer.NewClient(a.cfg.MailerAddr, a.cfg.MailerFromAddr, a.log)
	if err != nil {
		return fmt.Errorf("%s: failed to connect to mailer service: %w", operation, err)
	}

	sg := subscriber.NewRepo(db)
	fmter := formatter.NewWithDate()
	rw, err := ratewatcher.NewClient(a.cfg.RateWatcherAddr, a.log.With("source", "rateWatcher"))
	if err != nil {
		return fmt.Errorf("%s: failed to connect to rate watcher: %w", operation, err)
	}

	mailSender := sender.New(
		m,
		sg,
		fmter,
		rw,
		a.log.With("source", "cron sender"),
	)

	cronAdapter := sender.NewCronJobAdapter(mailSender, a.log.With("source", "adapter"))
	job := cron.NewDailyJob(cronJobHour, cronJobMinute, a.log.With("source", "cron"))
	job.Do(cronAdapter)

	l, err := net.Listen("tcp", net.JoinHostPort("", a.cfg.Port))
	if err != nil {
		return fmt.Errorf("%s: failed to start listener on port %s: %w", operation, a.cfg.Port, err)
	}

	return a.srv.Serve(l)
}

// GracefulStop method gracefully stop the server. It listens to the OS sigals.
// After it recieves signal it terminates all currently active servers,
// client, connections (if any) and gracefully exits.
func (a *App) GracefulStop() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	signal := <-ch
	a.log.Info("Recieved stop signal. Terminating...", "signal", signal)
	a.srv.Stop()
	a.log.Info("Successfully terminated server. Bye!")
}
