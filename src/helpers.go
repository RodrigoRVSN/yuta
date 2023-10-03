package main

import "github.com/aws/aws-lambda-go/events"

func OkFeedback(message string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: `{"message": "` + message + `"}`,
	}
}


func BadRequestFeedback(message string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: 400,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: `{"message": "` + message + `"}`,
	}
}

