package slack

import (
	"encoding/json"
	"net/url"
)

// APIGatewayRequest is a request from the AWS API Gateway when using the default Lambda proxy.
type APIGatewayRequest struct {
	Resource              string            `json:"resource"`
	Path                  string            `json:"path"`
	HTTPMethod            string            `json:"httpMethod"`
	Headers               map[string]string `json:"headers"`
	QueryStringParameters string            `json:"queryStringParameters"`
	PathParameters        string            `json:"pathParameters"`
	Body                  string            `json:"body"`
	ParsedBody            H                 `json:"-"`
}

// ParseBody parses the body and sets the ParsedBody attribute
func (request *APIGatewayRequest) ParseBody() error {
	if request.Headers["Content-Type"] == "application/x-www-form-urlencoded" {
		return request.parseFormBody()
	} else if request.Headers["Content-Type"] == "application/json" {
		return request.parseJSONBody()
	}

	return nil
}

func (request *APIGatewayRequest) parseJSONBody() error {
	return json.Unmarshal([]byte(request.Body), &request.ParsedBody)
}

func (request *APIGatewayRequest) parseFormBody() error {
	values, err := url.ParseQuery(request.Body)
	if err != nil {
		return err
	}

	for key, values := range values {
		request.ParsedBody[key] = values[0]
	}

	return nil
}
