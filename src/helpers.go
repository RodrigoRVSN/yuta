package main

import "github.com/aws/aws-lambda-go/events"

const (
	StatusOK         = 200
	StatusBadRequest = 400
)

func OkFeedback(message string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: `{"message": "` + message + `"}`,
	}
}

func BadRequestFeedback(message string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: StatusBadRequest,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: `{"message": "` + message + `"}`,
	}
}
