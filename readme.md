Sending a request to API Gateway triggers the Lambda function, which will upload the file to S3, already on CDN, and return the CDN link.

> Flow at Excalidraw

![image](https://github.com/RodrigoRVSN/yuta/assets/75763403/51200a3f-cb75-448b-9e93-654011d7d442)


> Request example

![image](https://github.com/RodrigoRVSN/yuta/assets/75763403/87c0f03f-e5e8-43e5-a519-8a823ff90a84)

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
