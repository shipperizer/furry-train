package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kelseyhightower/envconfig"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	checkGW "github.com/shipperizer/furry-train/check"
	"github.com/shipperizer/furry-train/health"
)

type EnvSpec struct {
	LogLevel         string `envconfig:"log_level"`
	HTTPPort         string `envconfig:"http_port"`
	GRPCheckEndpoint string `envconfig:"grpc_check_endpoint"`
}

// TODO @shipperizer refactor to follow https://github.com/grpc-ecosystem/grpc-gateway/blob/master/examples/internal/gateway/main.go
func main() {

	logger, err := zap.NewDevelopment()
	defer logger.Sync()

	if err != nil {
		panic(err.Error())
	}

	var specs EnvSpec
	err = envconfig.Process("", &specs)

	if err != nil {
		panic(err.Error())
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	logger.Sugar().Info(specs)

	err = checkGW.RegisterCheckServiceHandlerFromEndpoint(ctx, mux, specs.GRPCheckEndpoint, opts)

	if err != nil {
		logger.Sugar().Fatal(err)
	}

	api := http.NewServeMux()
	api.HandleFunc("/api/v1/status", health.Status())
	api.Handle("/api/v1/metrics", promhttp.Handler())
	api.Handle("/", mux)

	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%v", specs.HTTPPort),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      api, // Pass our instance of gorilla/mux in.
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Sugar().Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctxSrv, cancelSrv := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelSrv()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctxSrv)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	logger.Sugar().Info("Shutting down")
	os.Exit(0)
}
