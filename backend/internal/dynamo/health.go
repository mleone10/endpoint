package dynamo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Health returns true if the DynamoDB table is accessible, otherwise false.
func (c *Client) Health() bool {
	_, err := c.db.DescribeTable(&dynamodb.DescribeTableInput{TableName: aws.String(endpointTableName)})
	if err != nil {
		return false
	}

	return true
}
