package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ryus08/jiraTagger/apigw"
	"github.com/ryus08/jiraTagger/controller"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, req *events.APIGatewayProxyRequest) (*apigw.APIResponse, error) {
	receive := &controller.Receive{}

	requestBody := &controller.RequestBody{}

	if req.HTTPMethod == http.MethodGet {
		requestBody.Content = "Hello!"

	} else {
		fmt.Printf("%s\n", req.Body)
		//Unmarshaling request body
		err := json.Unmarshal([]byte(req.Body), requestBody)

		if err != nil {
			return apigw.Err(err)
		}
	}

	response := receive.Handler(requestBody)

	return apigw.ResponseWithHeaders(response, http.StatusOK, apigw.Headers{
		"Content-Type": "application/json",
	})
}

func main() {
	lambda.Start(Handler)
}
