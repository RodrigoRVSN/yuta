package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("🚀 Received event! Running...")

	awsRegion := "us-east-1"
	bucketName := "test-yuta-devto"

	fileHeaders, headersError := GetHeaders(event, 10485760)

	if headersError != nil {
		return BadRequestFeedback("Error reading headers: " + headersError.Error()), nil
	}

	fileName, uploadError := UploadFile(fileHeaders[0], awsRegion, bucketName)

	if uploadError != nil {
		return BadRequestFeedback("Error uploading file to S3: " + uploadError.Error()), nil
	}

	cloudFrontURL := fmt.Sprintf("%s/%s", "https://d3ijuim4ecbckg.cloudfront.net", fileName)

	return OkFeedback("✅ File uploaded successfully to CDN: " + cloudFrontURL), nil
}
