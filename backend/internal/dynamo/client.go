package dynamo

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	endpointTableName = "endpoint"
	endpointPK        = "pk"
	endpointSK        = "sk"
)

// Client represents a DynamoDB accessor.
type Client struct {
	db *dynamodb.DynamoDB
}

// NewClient returns an initialized Dynamo Client.
func NewClient() (*Client, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return &Client{
		db: dynamodb.New(sess),
	}, nil
}
