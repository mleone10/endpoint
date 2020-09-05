package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/handlerfunc"
	"github.com/mleone10/endpoint/internal/server"
)

var adapter *handlerfunc.HandlerFuncAdapter

func init() {
	adapter = handlerfunc.New(server.NewServer().ServeHTTP)
}

func serverHandler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return adapter.Proxy(req)
}

func main() {
	lambda.Start(serverHandler)
}
