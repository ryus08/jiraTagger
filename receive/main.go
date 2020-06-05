package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ryus08/jiraTagger/apigw"
)

type RequestBody struct {
	Content   string `json:"content"`
	Key       string `json:"key"`
	Token     string `json:"token"`
	Challenge string `json:"challenge"`
}

type ResponseBody struct {
	Message string

	*RequestBody
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, req *events.APIGatewayProxyRequest) (*apigw.APIResponse, error) {

	//Unmarshaling request body
	requestBody := &RequestBody{}
	err := json.Unmarshal([]byte(req.Body), requestBody)

	if err != nil {
		return apigw.Err(err)
	}

	//custom response headers
	responseHeaders := apigw.Headers{
		"Content-Type":           "application/json",
		"X-MyCompany-Func-Reply": "world-handler",
	}

	fmt.Printf("%s\n", req.Body)

	//echo back the same request payload with success message
	response := &ResponseBody{Message: "success", RequestBody: requestBody}
	return apigw.ResponseWithHeaders(response, http.StatusOK, responseHeaders)
}

func main() {
	lambda.Start(Handler)
}
