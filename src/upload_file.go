package main

import (
	"bytes"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type UploadRes struct {
	S3Path string
	S3URL  string
}

func UploadFile(fileHeader *multipart.FileHeader, region string, bucket string) (string, error) {
	file, openError := fileHeader.Open()

	if openError != nil {
		return "", openError
	}

	var fileContents bytes.Buffer
	_, readError := fileContents.ReadFrom(file)
	if readError != nil {
		return "", readError
	}

	awsSession, sessionError := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if sessionError != nil {
		return "", sessionError
	}

	uploader := s3manager.NewUploader(awsSession)

	fileName := fileHeader.Filename

	_, uploadError := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(fileContents.Bytes()),
	})

	if uploadError != nil {
		return "", uploadError
	}

	return fileName, nil
}
