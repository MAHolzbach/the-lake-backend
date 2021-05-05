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

	type ServiceData struct {
		ServiceName  string `json:"serviceName"`
		ServicePrice string `json:"servicePrice"`
	}

	services := []ServiceData{
		{
			ServiceName:  "Jetski",
			ServicePrice: "15.99",
		},
		{
			ServiceName:  "Waterski",
			ServicePrice: "25.99",
		},
		{
			ServiceName:  "Fishing",
			ServicePrice: "1.99",
		},
		{
			ServiceName:  "Kayak",
			ServicePrice: "4.99",
		},
		{
			ServiceName:  "Acuba",
			ServicePrice: "19.99",
		},
		{
			ServiceName:  "Canoe",
			ServicePrice: "5.99",
		},
	}

	body, err := json.Marshal(services)

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
			"X-TheLake-Func-Reply": "services-handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
