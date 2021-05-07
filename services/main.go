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
		ServiceName  string `json:"name"`
		RenderName   string `json:"renderName"`
		ServicePrice string `json:"price"`
	}

	services := []ServiceData{
		{
			ServiceName:  "jetski",
			RenderName:   "Jetskiing",
			ServicePrice: "15.99",
		},
		{
			ServiceName:  "waterski",
			RenderName:   "Waterskiing",
			ServicePrice: "25.99",
		},
		{
			ServiceName:  "windsurf",
			RenderName:   "Windsurfing",
			ServicePrice: "11.99",
		},
		{
			ServiceName:  "kayak",
			RenderName:   "Kayaking",
			ServicePrice: "4.99",
		},
		{
			ServiceName:  "scuba",
			RenderName:   "Scuba Diving",
			ServicePrice: "19.99",
		},
		{
			ServiceName:  "canoe",
			RenderName:   "Canoeing",
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
			"Content-Type":                     "application/json",
			"X-TheLake-Func-Reply":             "services-handler",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
