package dynamo

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/mleone10/endpoint/internal/account"
	"github.com/mleone10/endpoint/internal/station"
)

const skPrefixStation = "STATION"

type stationItem struct {
	itemKey
	Station station.Station `json:"station"`
}

// SaveStation persists a given account's station in Dynamo.
func (c *Client) SaveStation(id account.ID, s station.Station) error {
	if id == "" {
		return ErrorInvalidID
	}

	item, err := dynamodbattribute.MarshalMap(&stationItem{
		itemKey: itemKey{
			PK: id.String(),
			SK: fmt.Sprintf("%s#%s", skPrefixStation, s.ID),
		},
		Station: s,
	})
	if err != nil {
		return fmt.Errorf("failed to marshal station to dynamodb item: %v", err)
	}

	_, err = c.db.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(endpointTableName),
		Item:      item,
	})
	if err != nil {
		return fmt.Errorf("failed to persist station to DB: %v", err)
	}

	return nil
}
