package slack

import "encoding/json"

// CommandResponseType is a flag to select the way how the response is going to be presented to the user
type CommandResponseType string

const (
	// CommandResponseTypeEphemeral only displays the message to the user
	CommandResponseTypeEphemeral CommandResponseType = "ephemeral"

	// CommandResponseTypeInChannel sends the response to the channel
	CommandResponseTypeInChannel CommandResponseType = "in_channel"
)

// APIGatewayResponse is a response for the AWS API Gateway when using the default Lambda proxy.
type APIGatewayResponse struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}

// NewAPIGatewayJSONResponse creates a new response in JSON format
func NewAPIGatewayJSONResponse(code int, body H) *APIGatewayResponse {
	response := &APIGatewayResponse{
		StatusCode: code,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	data, err := json.Marshal(body)
	if err == nil {
		response.Body = string(data)
	}

	return response
}

// NewCommandResponse is a response for a slash command
func NewCommandResponse(responseType CommandResponseType, text string, attachments []*Attachment) *APIGatewayResponse {
	return NewAPIGatewayJSONResponse(200, H{
		"text":          text,
		"response_type": responseType,
		"attachments":   attachments,
	})
}
