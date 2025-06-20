package main

import (
    "context"
    "net/http"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    return events.APIGatewayProxyResponse{
        Body:       "pong",
        StatusCode: http.StatusOK,
    }, nil
}

func main() {
    lambda.Start(Handler)
}