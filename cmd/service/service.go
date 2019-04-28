package service

import (
	"flag"
	"fmt"
	"github.com/go-kit/kit/log/level"
	"github.com/lattecake/kpaas-alert/pkg/config"
	"github.com/lattecake/kpaas-alert/pkg/db"
	"net"
	http1 "net/http"
	"os"
	"os/signal"
	"syscall"

	endpoint1 "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/lattecake/kpaas-alert/pkg/endpoint"
	"github.com/lattecake/kpaas-alert/pkg/grpc"
	"github.com/lattecake/kpaas-alert/pkg/grpc/pb"
	"github.com/lattecake/kpaas-alert/pkg/http"
	"github.com/lattecake/kpaas-alert/pkg/service"
	"github.com/oklog/oklog/pkg/group"
	grpc1 "google.golang.org/grpc"
)

var logger log.Logger

var fs = flag.NewFlagSet("kpaas-alert", flag.ExitOnError)
var httpAddr = fs.String("http-addr", ":8080", "HTTP listen address")
var grpcAddr = fs.String("grpc-addr", ":8082", "gRPC listen address")
var configAddr = fs.String("config-addr", "./conf/app.yaml", "config file address")

func Run() {
	if err := fs.Parse(os.Args[1:]); err != nil {
		panic(err)
	}

	if err := config.InitConfigFile(*configAddr); err != nil {
		panic(err)
	}

	logger = log.NewLogfmtLogger(log.StdlibWriter{})
	logger = log.With(logger, "app", "kpaas-server")
	logger = log.With(logger, "caller", log.DefaultCaller)
	logger = level.NewFilter(logger, level.AllowInfo())

	if db.GetMysqlSession() == nil {
		_ = level.Error(logger).Log("db", "GetMysqlSession", "err", "db connect error")
		os.Exit(1)
	}

	svc := service.New(getServiceMiddleware(logger), logger)
	eps := endpoint.New(svc, getEndpointMiddleware(logger))
	g := createService(eps)
	initCancelInterrupt(g)
	_ = logger.Log("exit", g.Run())

}
func initGRPCHandler(endpoints endpoint.Endpoints, g *group.Group) {
	options := defaultGRPCOptions(logger)

	grpcServer := grpc.NewGRPCServer(endpoints, options)
	grpcListener, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		_ = logger.Log("transport", "gRPC", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		_ = logger.Log("transport", "gRPC", "addr", *grpcAddr)
		baseServer := grpc1.NewServer()
		pb.RegisterKpaasAlertServer(baseServer, grpcServer)
		return baseServer.Serve(grpcListener)
	}, func(error) {
		if err = grpcListener.Close(); err != nil {
			panic(err)
		}
	})

}
func getServiceMiddleware(logger log.Logger) (mw []service.Middleware) {
	mw = []service.Middleware{}
	mw = addDefaultServiceMiddleware(logger, mw)
	mw = append(mw, service.AuthorizationMiddleware(logger))

	return
}
func getEndpointMiddleware(logger log.Logger) (mw map[string][]endpoint1.Middleware) {
	mw = map[string][]endpoint1.Middleware{}

	addDefaultEndpointMiddleware(logger, mw)

	return
}

func initCancelInterrupt(g *group.Group) {
	cancelInterrupt := make(chan struct{})
	g.Add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-cancelInterrupt:
			return nil
		}
	}, func(error) {
		close(cancelInterrupt)
	})
}

func initHttpHandler(endpoints endpoint.Endpoints, g *group.Group) {
	options := defaultHttpOptions(logger)

	httpHandler := http.NewHTTPHandler(endpoints, options)
	httpListener, err := net.Listen("tcp", *httpAddr)
	if err != nil {
		_ = logger.Log("transport", "HTTP", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		_ = logger.Log("transport", "HTTP", "addr", *httpAddr)
		return http1.Serve(httpListener, httpHandler)
	}, func(error) {
		if err = httpListener.Close(); err != nil {
			panic(err)
		}
	})

}
