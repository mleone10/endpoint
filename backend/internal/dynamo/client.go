package dynamo

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	endpointTableName = "endpoint"
	endpointPK        = "pk"
	endpointSK        = "sk"
)

var (
	// ErrorItemNotFound is returned when an expected item is not found in the database.  Callers can use this to return an HTTP 404 to clients.
	ErrorItemNotFound = fmt.Errorf("item not found")
	// ErrorConflict is returned when an attempted write fails due to a condition or conflict with an existing resource.
	ErrorConflict = fmt.Errorf("conflict with existing resource")
)

// Client represents a DynamoDB accessor.
type Client struct {
	db *dynamodb.DynamoDB
}

type itemKey struct {
	PK string `json:"pk"`
	SK string `json:"sk"`
}

// NewClient returns an initialized Dynamo Client.
func NewClient() *Client {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return &Client{
		db: dynamodb.New(sess),
	}
}
