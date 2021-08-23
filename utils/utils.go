package utils

import (
"github.com/aws/aws-lambda-go/events"
)

// Creates an ApiGatewayProxyResponse with CORS headers based on the given status code and marshalled json body
func CreateResponse(status int, body string) (events.APIGatewayProxyResponse, error) {

	// Cloudflare - We support the GET, POST, HEAD, and OPTIONS methods from any origin,// and allow any header on requests. These headers must be present// on all responses to all CORS preflight requests. In practice, this means// all responses to OPTIONS requests.
	response := events.APIGatewayProxyResponse{
		StatusCode: status,
		Headers: map[string]string{
			"x-custom-header":              "*",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "GET,HEAD,POST,OPTIONS",
			"Access-Control-Max-Age":       "86400",
		},
		Body: body,
	}

	return response, nil
}
