package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ryus08/jiraTagger/apigw"
	"github.com/ryus08/jiraTagger/controller"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, req *events.APIGatewayProxyRequest) (*apigw.APIResponse, error) {
	receive := &controller.Receive{}

	if req.HTTPMethod == http.MethodGet {
		req.Body = "{Content: \"Hello!\"}"
	}

	doubleHeaders := http.Header{}

	for index, element := range req.Headers {
		doubleHeaders[index] = []string{element}
	}

	e := receive.Authorize(doubleHeaders, &req.Body)
	var response interface{}
	var statusCode int
	if e == nil {
		response, e = receive.Handler(&req.Body)
		statusCode = http.StatusOK
	}

	if e != nil {
		response = e
		statusCode = http.StatusInternalServerError
	}

	return apigw.ResponseWithHeaders(response, statusCode, apigw.Headers{
		"Content-Type": "application/json",
	})
}

func main() {
	lambda.Start(Handler)
}
