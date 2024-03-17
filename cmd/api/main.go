package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	version           = "1.0.0"
	portStr           = "port"
	envStr            = "env"
	defaultPort       = 4000
	defaultEnv        = "development"
	defaultEnvMessage = "application environment {development|staging|production}"
	serverMessage     = "server port to listen on"
	errMessage        = "the server encountered a problem and could not process your request"
	notFoundMessage   = "the requested resource could not be found"
	methodMessage     = "the %s method is not supported"
)

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, portStr, defaultPort, serverMessage)
	flag.StringVar(&cfg.env, envStr, defaultEnv, defaultEnvMessage)
	flag.Parse()

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
