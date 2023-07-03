package server

import (
	"context"
	"net/http"
	"recipes/internal/handler"
	"recipes/internal/router"
	"recipes/pkg/logger"
	"time"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	srv *http.Server
	lg  logger.Logger
}

func New(lg logger.Logger, addr string, h *handler.Handler) *Server {
	return &Server{
		srv: &http.Server{
			Handler: router.New(h),
			Addr:    addr,
		},
		lg: lg,
	}
}

func (s *Server) Start(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(s.srv.ListenAndServe)
	eg.Go(func() error {
		s.lg.Infoln("Recipe service started. App addr:", s.srv.Addr)
		<-ctx.Done()
		s.lg.Infoln("Recipe service sthutting down")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := s.srv.Shutdown(shutdownCtx)
		if err != nil {
			s.lg.Errorln(err.Error())
		}
		return err
	})
	return eg.Wait()
}
