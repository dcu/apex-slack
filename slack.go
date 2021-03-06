package slack

import (
	"encoding/json"
	"os"

	apex "github.com/apex/go-apex"
)

var (
	_commandHandlers   = map[string]CommandHandler{}
	_verificationToken = os.Getenv("SLACK_VERIFICATION_TOKEN")
)

// H is the representation of a generic hash
type H map[string]interface{}

func (h H) String(key string) string {
	value, ok := h[key].(string)
	if !ok {
		return ""
	}

	return value
}

// CommandHandler is a function that handles a command
type CommandHandler func(request *APIGatewayRequest) *APIGatewayResponse

// OnCommand receives the name of the command and a handler
func OnCommand(name string, handler CommandHandler) {
	_commandHandlers[name] = handler
}

// Handle handles the incoming apex event
func Handle(message json.RawMessage, ctx *apex.Context) (interface{}, error) {
	request := &APIGatewayRequest{
		ParsedBody: H{},
	}
	err := json.Unmarshal(message, &request)
	if err != nil {
		return NewAPIGatewayJSONResponse(400, H{"text": "Invalid request from API Gateway"}), nil
	}

	err = request.ParseBody()
	if err != nil {
		return NewAPIGatewayJSONResponse(400, H{"text": "Invalid body format"}), nil
	}

	if _verificationToken != "" && _verificationToken != request.ParsedBody.String("token") {
		return NewAPIGatewayJSONResponse(400, H{"text": "Invalid verification token"}), nil
	}

	if handler, ok := _commandHandlers[request.ParsedBody.String("command")]; ok {
		return handler(request), nil
	}

	return NewAPIGatewayJSONResponse(400, H{"text": "Unable to find a handler for your request"}), nil
}

// Init initializes the slack handler
func Init() {
	apex.HandleFunc(Handle)
}
