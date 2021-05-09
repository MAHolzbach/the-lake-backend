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

	type PintaData struct {
		BoatName   string `json:"boatName"`
		RenderName string `json:"renderName"`
		Class      string `json:"class"`
		Width      string `json:"width"`
		Length     string `json:"length"`
		Capacity   int    `json:"capacity"`
		Price      string `json:"price"`
	}

	pintaData := PintaData{
		BoatName:   "pinta",
		RenderName: "The Pinta",
		Class:      "Displacement Yacht",
		Width:      "10 meters",
		Length:     "24 meters",
		Capacity:   16,
		Price:      "350",
	}

	body, err := json.Marshal(pintaData)

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
			"X-TheLake-Func-Reply":             "pinta-handler",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
