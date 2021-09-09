package main

import "go.uber.org/zap"

type api struct {
	Log *zap.SugaredLogger;
}

func newApplication() (*api, error) {
	z, _ := zap.NewProduction()

	a := &api {
		Log: z.Sugar(),
	}

	return a, nil
}
