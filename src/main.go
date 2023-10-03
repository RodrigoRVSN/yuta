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
		return BadRequestFeedback("Error reading headers: " + headersError.Error()), nil
	}

	fileName := uuid.New().String()
	fileExtension := path.Ext(fileHeaders[0].Filename)
	fileName = fileName + fileExtension

	_, uploadError := UploadFile(fileHeaders[0], awsRegion, bucketName, fileName)

	if uploadError != nil {
		return BadRequestFeedback("Error uploading file to S3: " + uploadError.Error()), nil
	}

	cloudFrontURL := fmt.Sprintf("%s/%s", "https://d1dlccvqvhrnl1.cloudfront.net", fileName)

	return OkFeedback("✅ File uploaded successfully to CDN: " + cloudFrontURL), nil
}
