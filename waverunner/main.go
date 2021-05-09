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

	type WindrunnerData struct {
		BoatName   string `json:"boatName"`
		RenderName string `json:"renderName"`
		Class      string `json:"class"`
		Width      string `json:"width"`
		Length     string `json:"length"`
		Capacity   int    `json:"capacity"`
		Price      string `json:"price"`
	}

	windrunnerData := WindrunnerData{
		BoatName:   "windrunner",
		RenderName: "The Windrunner",
		Class:      "Speed Boat",
		Width:      "2.5 meters",
		Length:     "10 meters",
		Capacity:   2,
		Price:      "125",
	}

	body, err := json.Marshal(windrunnerData)

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
			"X-TheLake-Func-Reply":             "windrunner-handler",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
