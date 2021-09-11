package main

import (
	"go.uber.org/zap"
	"log"
)

type config struct {
	port int
	env  string
	db struct {
		dsn string
	}
}

type api struct {
	zLogger *zap.SugaredLogger;
}

func newApplication() (*api, error) {
	zLog := newLogger()
	
	a := &api {
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
