package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func hello() (events.APIGatewayProxyResponse, error) {
	body := "Hello World"

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       body,
	}, nil
}
func main() {
	lambda.Start(hello)
}

// COMMANDS
/*
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o bootstrap cmd/main.go

zip -9 -j bootstrap.zip bootstrap

npm init

sudo npm install -g serverless

serverless

sls config credentials --provider aws --key <aws access key> --secret <aws secret key>

sls deploy
*/
