package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Question struct {
	Question      string `json:"Question"`
	NumberToGuess string `json:"NumberToGuess"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	language := request.QueryStringParameters["language"]
	if language == "" {
		language = "en-US"
	}

	// seed := request.QueryStringParameters["seed"]

	question := &Question{
		Question:      "test",
		NumberToGuess: "132",
	}

	response, err := json.Marshal(question)
	if err != nil {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(response),
	}, nil
}

func serverError(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
