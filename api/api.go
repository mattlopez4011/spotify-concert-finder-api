package main

import (
	"flag"
	"log"
	"os"

	"go.uber.org/zap"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type api struct {
	config config
	zLogger *zap.SugaredLogger
}

func newApi() (*api, error) {
	var cfg config
	flag.IntVar(&cfg.port, "port", 8080, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development | production ")
	flag.Parse()
	
	zLog := newLogger()
	
	a := &api {
		config: cfg,
		zLogger: zLog,
	}

	return a, nil
}

func newLogger() (*zap.SugaredLogger) {
	z, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}

	return z.Sugar()
}
