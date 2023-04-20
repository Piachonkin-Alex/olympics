package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"golang.org/x/sync/errgroup"
)

type Config struct {
	Listen       string        `yaml:"listen"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
	Profile      bool          `yaml:"profile"`
}

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type Server struct {
	httpServer *http.Server
	router     *chi.Mux
	cfg        *Config
	done       chan struct{}
}

func NewServer(cfg *Config) *Server {
	s := &Server{cfg: cfg, router: chi.NewRouter(), done: make(chan struct{})}
	s.httpServer = &http.Server{
		Addr:         cfg.Listen,
		Handler:      s.router,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	return s
}

func (s *Server) Start(ctx context.Context) error {
	if s.cfg.Profile {
		s.router.Get("/pprof/profile", pprof.Profile)
		s.router.Get("/pprof/trace", pprof.Trace)
		s.router.Get("/pprof/index", pprof.Index)
		s.router.Get("/pprof/heap", pprof.Handler("heap").ServeHTTP)
		s.router.Get("/pprof/goroutine", pprof.Handler("goroutine").ServeHTTP)
		s.router.Get("/pprof/threadcreate", pprof.Handler("threadcreate").ServeHTTP)
		s.router.Get("/pprof/allocs", pprof.Handler("allocs").ServeHTTP)
		s.router.Get("/pprof/block", pprof.Handler("block").ServeHTTP)
		s.router.Get("/pprof/mutex", pprof.Handler("mutex").ServeHTTP)
	}
	s.router.Get("/ping", s.getPing)

	stop := make(chan os.Signal, 2)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	gr, ctx := errgroup.WithContext(ctx)
	gr.Go(func() error {
		select {
		case <-ctx.Done():
		case <-stop:
			log.Println("Server shutdowned by signal")
		}
		s.done <- struct{}{}
		return nil
	})

	gr.Go(func() error {
		err := s.httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})

	return gr.Wait()
}

func (s *Server) Stop(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) Done() chan struct{} {
	return s.done
}

func (s *Server) getPing(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("pong"))
}

func (s *Server) AddMiddleware(mw func(http.Handler) http.Handler) {
	s.router.Use(mw)
}

func (s *Server) AddRoute(method, path string, fn http.Handler) {
	s.router.Method(method, path, fn)
}

func (s *Server) AddRoutes(routes []Route) {
	for _, route := range routes {
		s.router.Method(route.Method, route.Path, route.Handler)
	}
}

func (s *Server) SetNotFoundHandler(fn http.HandlerFunc) {
	s.router.NotFound(fn)
}
