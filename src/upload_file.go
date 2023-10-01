package main

import (
	"bytes"
	"mime/multipart"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type UploadRes struct {
	S3Path string
	S3URL  string
}

func UploadFile(fileHeader *multipart.FileHeader, region, bucket, fileName string) (*UploadRes, error) {
	file, openError := fileHeader.Open()

	if openError != nil {
		return nil, openError
	}

	var fileContents bytes.Buffer
	_, readError := fileContents.ReadFrom(file)
	if readError != nil {
		return nil, readError
	}

	awsSession, sessionError := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	if sessionError != nil {
		return nil, sessionError
	}

	uploader := s3manager.NewUploader(awsSession)

	uploadOutput, uploadError := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(fileContents.Bytes()),
	})

	if uploadError != nil {
		return nil, uploadError
	}

	return &UploadRes{
		S3Path: filepath.Join(bucket, fileName),
		S3URL:  uploadOutput.Location,
	}, nil
}
