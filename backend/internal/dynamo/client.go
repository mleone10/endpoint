package dynamo

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

/*
	Table schema (PK: userID:string, SK: sortKey:string)
	API-Key schema (SK: KEY#key:string, Attrs: apiKey:string, name:string, readOnly:string)
	GSI (PK: userID:string, SK: apiKey:string)
*/

const (
	endpointTableName = "endpoint"
	endpointPK        = "userID"
	endpointSK        = "sortKey"
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
