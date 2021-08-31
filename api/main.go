package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	chiAdapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
)

type config struct {
	port int
	env  string
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development | production ")
	flag.Parse()
	
	if cfg.env == "prod" {
		lambda.Start(handler)
	} else {
		err :=  http.ListenAndServe(":3000", r)
		if err != nil {
		// a.Log.Errorw("Error serving request", "error", err.Error())
	}
	}
}


// func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// 	chiLambda := chiAdapter.New(newRouter(nil))

// 	response, err := chiLambda.ProxyWithContext(ctx, req)
// 	if err != nil {
// 		return response, err
// 	}

// 	response.Headers = map[string]string{
// 		"x-custom-header":              "*",
// 		"Access-Control-Allow-Origin":  "*",
// 		"Access-Control-Allow-Methods": "GET,HEAD,POST,OPTIONS",
// 		"Access-Control-Max-Age":       "86400",
// 	}

// 	return response, nil
// }
