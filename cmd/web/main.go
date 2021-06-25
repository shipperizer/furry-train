package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type EnvSpec struct {
	Port     string `envconfig:"http_port"`
	LogLevel string `envconfig:"log_level"`
}

func Status(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Purple Bro"})
}

func main() {
	var specs EnvSpec
	err := envconfig.Process("", &specs)

	if err != nil {
		log.Fatal(err.Error())
	}

	level, err := log.ParseLevel(specs.LogLevel)

	if err != nil {
		level = log.DebugLevel
	}

	log.SetLevel(level)

	router := mux.NewRouter()
	router.HandleFunc("/api/v0/status", Status).Methods(http.MethodGet, http.MethodOptions)

	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%v", specs.Port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, // Pass our instance of gorilla/mux in.
	}

	go func() {
		log.Infof("Starting up on %v", specs.Port)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
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

	log.Println("Shutting down")
	os.Exit(0)
}
