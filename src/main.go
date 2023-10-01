package main

import (
	"context"
	"fmt"
	"path"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("🚀 Received event! Running...")

	awsRegion := "us-east-1"
	bucketName := "yuta-okkotsu"

	fileHeaders, headersError := GetHeaders(event, 10485760)

	if headersError != nil {
		fmt.Println("Failed to upload file to S3 reading headers:", headersError)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: `{"message": "Error uploading file to S3"}`,
		}, nil
	}

	fileName := uuid.New().String()
	fileExtension := path.Ext(fileHeaders[0].Filename)
	fileName = fileName + fileExtension

	_, uploadError := UploadFile(fileHeaders[0], awsRegion, bucketName, fileName)

	if uploadError != nil {
		fmt.Println("Failed to upload file to S3 uploading file:", uploadError)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: `{"message": "Error uploading file to S3"}`,
		}, nil
	}

	cloudFrontURL := fmt.Sprintf("%s/%s", "https://d1dlccvqvhrnl1.cloudfront.net", fileName)

	fmt.Println("✅ File uploaded successfully. CloudFront URL:", cloudFrontURL)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: `{"message": "File uploaded successfully"}`,
	}, nil
}
