package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/lattecake/kpaas-alert/pkg/db"
	"github.com/pkg/errors"
)

// KpaasAlertService describes the service.
type KpaasAlertService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Alert(ctx context.Context, data string) (rs string, err error)
	GetAlert(ctx context.Context, id int64) (rs string, err error)
}

type basicKpaasAlertService struct {
	logger log.Logger
}

func (b *basicKpaasAlertService) Alert(ctx context.Context, data string) (rs string, err error) {
	//_ = level.Info(b.logger).Log("data", data)
	prom, err := NewPrometheusAlerts([]byte(data))
	if err != nil {
		return
	}

	prom.Get()

	var ns, name string

	_ = level.Info(b.logger).Log("ns", ns, "name", name)

	if !db.Create(&db.Alert{
		Name:      name,
		Namespace: ns,
		Body:      prom.String(),
	}) {
		return "", errors.New("db create is false.")
	}

	return
}
func (b *basicKpaasAlertService) GetAlert(ctx context.Context, id int64) (rs string, err error) {
	// TODO implement the business logic of GetAlert
	return rs, err
}

// NewBasicKpaasAlertService returns a naive, stateless implementation of KpaasAlertService.
func NewBasicKpaasAlertService(logger log.Logger) KpaasAlertService {
	return &basicKpaasAlertService{logger: logger}
}

// New returns a KpaasAlertService with all of the expected middleware wired in.
func New(middleware []Middleware, logger log.Logger) KpaasAlertService {
	var svc KpaasAlertService = NewBasicKpaasAlertService(logger)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
