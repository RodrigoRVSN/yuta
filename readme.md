Sending a request to API Gateway triggers the Lambda function, which will upload the file to S3, already on CDN, and return the CDN link.

![image](https://github.com/RodrigoRVSN/yuta/assets/75763403/0820a281-e9eb-4c13-9437-4347d13f3819)

# Technologies

- Golang
- AWS SDK
- AWS S3
- AWS CloudFront
- AWS Lambda
- AWS API Gateway
- AWS CloudWatch

# Using AWS SAM CLI for testing lambda functions locally

```bash
go build -o yuta ./src/* && sam local start-api
```

or

```bash
scripts/run-sam.sh
```

> Curiosity: Yuta from JJK was chosen as the name of the repo because it storage a lot of cursed energy, it's a reference to the storage on S3
