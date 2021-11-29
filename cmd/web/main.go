package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/shipperizer/miniature-monkey/config"
	"github.com/shipperizer/miniature-monkey/core"
	"go.uber.org/zap"

	"github.com/shipperizer/furry-train/check"
	"github.com/shipperizer/furry-train/monitoring"
)

type EnvSpec struct {
	Port     string `envconfig:"http_port"`
	LogLevel string `envconfig:"log_level"`
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

	apiCfg := config.NewAPIConfig(
		"furry-train-web",
		nil,
		monitoring.NewPrometheusMonitor(monitoring.PrometheusConfig{Service: "furry-train-web", Logger: logger.Sugar()}),
		logger.Sugar(),
	)

	api := core.NewAPI(apiCfg)

	api.RegisterBlueprints(api.Router(), check.NewBlueprint(logger.Sugar()))

	srv := &http.Server{
		Addr:         "0.0.0.0:8000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      api.Handler(),
	}

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
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
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
