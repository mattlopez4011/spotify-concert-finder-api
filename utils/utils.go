package utils

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"

	
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

const numRandomBytes = 32

func GenerateRandomString() (string, error) {
	b, err := GenerateRandomBytes(numRandomBytes)
	if err != nil {
		return "", err
	}

	return EncodeBase64WithoutPadding(b), nil
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)

	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func EncodeBase64WithoutPadding(input []byte) string {
	encoded := base64.URLEncoding.EncodeToString(input)
	parts := strings.Split(encoded, "=")
	encoded = parts[0]

	return encoded
}

func (a *api) writeJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})
	wrapper[wrap] = data

	js, err := json.Marshal(wrapper)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (a *api) errorJSON(w http.ResponseWriter, err error) {
	type jsonError struct {
		Message string `json:"message"`
	}

	theError := jsonError{
		Message: err.Error(),
	}

	a.writeJSON(w, http.StatusBadRequest, theError, "error")

}
