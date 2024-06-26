package app

import (
	"log"
	"net/http"
	"time"
)

const (
	idleTimeout        = 240 * time.Second
	writeHeaderTimeout = 15 * time.Second
	readHeaderTimeout  = 30 * time.Second
)

// newServer constructs built-in standard *http.Server with some values
// set by default.
func newServer(h http.Handler, addr string, log *log.Logger) *http.Server {
	srv := &http.Server{
		Handler:      h,
		ErrorLog:     log,
		Addr:         addr,
		IdleTimeout:  idleTimeout,
		WriteTimeout: writeHeaderTimeout,
		ReadTimeout:  readHeaderTimeout,
	}
	return srv
}
