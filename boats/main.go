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

	type BoatData struct {
		BoatName   string `json:"name"`
		RenderName string `json:"renderName"`
		BoatPrice  string `json:"price"`
		Status     string `json:"status"`
	}

	boats := []BoatData{
		{
			BoatName:   "breeze",
			RenderName: "The Breeze",
			BoatPrice:  "150",
			Status:     "Popular",
		},
		{
			BoatName:   "lancer",
			RenderName: "The Lancer",
			BoatPrice:  "75",
			Status:     "",
		},
		{
			BoatName:   "nina",
			RenderName: "The Nina",
			BoatPrice:  "100",
			Status:     "Unavailable",
		},
		{
			BoatName:   "pinta",
			RenderName: "The Pinta",
			BoatPrice:  "350",
			Status:     "",
		},
		{
			BoatName:   "santaMaria",
			RenderName: "The Santa Maria",
			BoatPrice:  "80",
			Status:     "Budget",
		},
		{
			BoatName:   "waverunner",
			RenderName: "The Waverunner",
			BoatPrice:  "125",
			Status:     "Fast!",
		},
	}

	body, err := json.Marshal(boats)

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
			"X-TheLake-Func-Reply":             "boats-handler",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
