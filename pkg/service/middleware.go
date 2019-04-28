package service

import (
	"context"
	"github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(KpaasAlertService) KpaasAlertService

type loggingMiddleware struct {
	logger log.Logger
	next   KpaasAlertService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a KpaasAlertService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next KpaasAlertService) KpaasAlertService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Alert(ctx context.Context, data string) (rs string, err error) {
	defer func() {
		_ = l.logger.Log("method", "Alert", "s", data, "rs", rs, "err", err)
	}()
	return l.next.Alert(ctx, data)
}
func (l loggingMiddleware) GetAlert(ctx context.Context, id int64) (rs string, err error) {
	defer func() {
		_ = l.logger.Log("method", "GetAlert", "id", id, "rs", rs, "err", err)
	}()
	return l.next.GetAlert(ctx, id)
}

type authorizationMiddleware struct {
	logger log.Logger
	next   KpaasAlertService
}

func AuthorizationMiddleware(logger log.Logger) Middleware {
	return func(next KpaasAlertService) KpaasAlertService {
		return &authorizationMiddleware{logger, next}
	}
}

func (a authorizationMiddleware) Alert(ctx context.Context, data string) (rs string, err error) {

	return a.next.Alert(ctx, data)
}

func (a authorizationMiddleware) GetAlert(ctx context.Context, id int64) (rs string, err error) {
	return a.next.GetAlert(ctx, id)
}
