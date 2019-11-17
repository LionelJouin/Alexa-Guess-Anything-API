package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	language := request.QueryStringParameters["language"]
	if language == "" {
		language = "en-US"
	}

	// seed := request.QueryStringParameters["seed"]

	message := fmt.Sprintf(" { \"Message\" : \"Hello %s \" } ", language)

	return events.APIGatewayProxyResponse{Body: message, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
