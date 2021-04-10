package main

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {
	var buf bytes.Buffer

	type JetskiData struct {
		VehicleType string   `json:"vehicleType"`
		Price       string   `json:"price"`
		Colors      []string `json:"colors"`
	}

	jetskiData := JetskiData{
		VehicleType: "jetski",
		Price:       "24.99",
		Colors:      []string{"Red", "Blue", "Green", "Yellow"},
	}

	body, err := json.Marshal(jetskiData)

	if err != nil {
		return Response{StatusCode: 404}, err
	}

	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":         "application/json",
			"X-TheLake-Func-Reply": "jetski-handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}