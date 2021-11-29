package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go.uber.org/zap"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/shipperizer/furry-train/check"
	"github.com/shipperizer/furry-train/health"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type EnvSpec struct {
	HTTPPort string `envconfig:"http_port"`
	GRPCPort string `envconfig:"grpc_port"`
}

func NewgRPCServer() *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	}

	// plugin prometheus monitoring
	server := grpc.NewServer(opts...)

	health.NewRPC().Register(server)
	check.NewRPC().Register(server)

	// TODO @shipperizer this increases cardinality of metric, use an env var for enabling it
	grpc_prometheus.EnableHandlingTimeHistogram()
	// init prometheus metrics provided by grpc_prometheus
	grpc_prometheus.Register(server)

	// Register reflection service on gRPC server
	reflection.Register(server)

	return server
}

func main() {
	logger, err := zap.NewDevelopment()
	defer logger.Sync()

	if err != nil {
		panic(err.Error())
	}

	var specs EnvSpec
	err = envconfig.Process("", &specs)

	if err != nil {
		logger.Sugar().Fatal(err.Error())
	}

	// AWS NLB TLS termination.
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", specs.GRPCPort))

	if err != nil {
		logger.Sugar().Fatal("Error: ", err)
		return
	}

	monitorAPI := mux.NewRouter()
	monitorAPI.PathPrefix("/api/v1/metrics").Handler(promhttp.Handler()).Name("v1.metrics")

	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%v", specs.HTTPPort),
		// Good practice to set timeouts to avoid Slowloris attacks.
		// WriteTimeout: time.Second * 15,
		ReadTimeout: time.Second * 15,
		IdleTimeout: time.Second * 60,
		Handler:     monitorAPI,
	}

	api := NewgRPCServer()

	go func() {
		logger.Sugar().Infof("Starting GRPC Server up on %v", specs.GRPCPort)
		if err := api.Serve(listener); err != nil {
			logger.Sugar().Fatal(err)
		}
	}()

	go func() {
		logger.Sugar().Infof("Starting up on %v", specs.HTTPPort)
		if err := srv.ListenAndServe(); err != nil {
			logger.Sugar().Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)

	// Block until we receive our signal.
	<-c

	api.GracefulStop()

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	logger.Sugar().Info("Shutting down")
	os.Exit(0)
}
