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

// ListStations returns a list of all stations for a given account.
func (c *Client) ListStations(id account.ID) ([]station.ID, error) {
	uidKey, skPrefixKey := ":uid", ":station"
	if id == "" {
		return nil, ErrorInvalidID
	}

	res, err := c.db.Query(&dynamodb.QueryInput{
		TableName:              aws.String(endpointTableName),
		KeyConditionExpression: aws.String(fmt.Sprintf("%s = %s and begins_with(%s, %s)", endpointPK, uidKey, endpointSK, skPrefixKey)),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			uidKey:      {S: aws.String(id.String())},
			skPrefixKey: {S: aws.String(skPrefixStation)},
		},
	})
	if err != nil {
		return nil, ErrorInternalServerError
	}

	ids := []station.ID{}
	for _, i := range res.Items {
		var s stationItem
		dynamodbattribute.UnmarshalMap(i, &s)
		ids = append(ids, s.Station.ID)
	}

	return ids, nil
}

// GetStation retrieves a given account's station from Dynamo.
func (c *Client) GetStation(uid account.ID, sid station.ID) (station.Station, error) {
	res, err := c.db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(endpointTableName),
		Key: map[string]*dynamodb.AttributeValue{
			endpointPK: {S: aws.String(uid.String())},
			endpointSK: {S: aws.String(fmt.Sprintf("%s#%s", skPrefixStation, sid))},
		},
	})
	if err != nil {
		return station.Station{}, fmt.Errorf("failed to get station: %v", err)
	}

	if res.Item == nil {
		return station.Station{}, ErrorItemNotFound
	}

	var s stationItem
	err = dynamodbattribute.UnmarshalMap(res.Item, &s)
	if err != nil {
		return station.Station{}, fmt.Errorf("failed to unmarshal station response: %v", err)
	}

	return s.Station, nil
}

// DeleteStation deletes a given account's station from Dynamo.
func (c *Client) DeleteStation(uid account.ID, sid station.ID) error {
	_, err := c.db.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(endpointTableName),
		Key: map[string]*dynamodb.AttributeValue{
			endpointPK: {S: aws.String(uid.String())},
			endpointSK: {S: aws.String(fmt.Sprintf("%s#%s", skPrefixStation, sid))},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to delete station [%v]: %v", sid, err)
	}

	return nil
}
