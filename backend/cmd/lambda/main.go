package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/handlerfunc"
	"github.com/mleone10/endpoint/internal/api"
	"github.com/mleone10/endpoint/internal/auth"
	"github.com/mleone10/endpoint/internal/dynamo"
)

var adapter *handlerfunc.HandlerFuncAdapter
var db *dynamo.Client

func init() {
	db := dynamo.NewClient()
	authr, err := auth.NewAuthenticator()
	if err != nil {
		log.Panicf("Failed to initialize authenticator: %v", err)
	}
	adapter = handlerfunc.New(api.NewServer(db, authr).ServeHTTP)
}

func serverHandler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return adapter.Proxy(req)
}

func main() {
	lambda.Start(serverHandler)
}
