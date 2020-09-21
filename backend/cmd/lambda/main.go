package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/handlerfunc"
	"github.com/mleone10/endpoint/internal/dynamo"
	"github.com/mleone10/endpoint/internal/server"
)

var adapter *handlerfunc.HandlerFuncAdapter
var db *dynamo.Client

func init() {
	db := dynamo.NewClient()
	adapter = handlerfunc.New(server.NewServer(db).ServeHTTP)
}

func serverHandler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return adapter.Proxy(req)
}

func main() {
	lambda.Start(serverHandler)
}
