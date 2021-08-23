package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	chiAdapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	chiLambda := chiAdapter.New(newRouter(nil))

	response, err := chiLambda.ProxyWithContext(ctx, req)
	if err != nil {
		return response, err
	}

	response.Headers = map[string]string{
		"x-custom-header":              "*",
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "GET,HEAD,POST,OPTIONS",
		"Access-Control-Max-Age":       "86400",
	}

	return response, nil
}
