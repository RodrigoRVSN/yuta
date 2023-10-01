package main

import (
	"encoding/base64"
	"fmt"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

func GetHeaders(lambdaReq events.APIGatewayProxyRequest, maxFileSizeBytes int64) ([]*multipart.FileHeader, error) {
	headers := http.Header{}

	for header, values := range lambdaReq.Headers {
		headers.Add(header, values)
	}

	contentType := headers.Get("Content-Type")
	if contentType == "" {
		return nil, fmt.Errorf("Content-Type header not found")
	}

	_, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		return nil, err
	}

	boundary := params["boundary"]
	if boundary == "" {
		return nil, fmt.Errorf("boundary not found in Content-Type header")
	}

	stringReader := strings.NewReader(lambdaReq.Body)
	b64Reader := base64.NewDecoder(base64.StdEncoding, stringReader)
	multipartReader := multipart.NewReader(b64Reader, boundary)

	form, err := multipartReader.ReadForm(maxFileSizeBytes)
	if err != nil {
		return nil, err
	}

	var files []*multipart.FileHeader

	for _, currentFileHeaders := range form.File {
		files = append(files, currentFileHeaders...)
	}

	return files, nil
}
