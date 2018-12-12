package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type resp struct {
	Links []link `json:"links"`
}

type link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

func handler(
	ctx context.Context,
	request events.APIGatewayProxyRequest,
) (
	events.APIGatewayProxyResponse,
	error,
) {
	apiURL := os.Getenv("API_URL")
	var resp resp
	resp.Links = []link{
		{
			Rel:  "self",
			Href: fmt.Sprintf("%s", apiURL),
		},
		{
			Rel:  "monkeys",
			Href: fmt.Sprintf("%s/monkeys", apiURL),
		},
	}

	body, _ := json.Marshal(resp)
	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
