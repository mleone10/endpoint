package dynamo

import "github.com/mleone10/endpoint/internal/account"

const (
	skPrefixUser   = "USER"
	skPrefixAPIKey = "KEY"
)

// User represents a user's basic info in the database
type User struct {
	itemKey
}

// APIKey represents an API key database item
type APIKey struct {
	itemKey
	Key      string `json:"key"`
	ReadOnly bool   `json:"readOnly"`
}

// SaveAPIKey upserts an APIKey in the Dynamo database
func (c *Client) SaveAPIKey(apiKey *account.APIKey) error {
	return nil
}

// ListAPIKeys returns a slice of APIKeys for a given account ID
func (c *Client) ListAPIKeys(account.ID) ([]account.APIKey, error) {
	return nil, nil
}

// // SaveUser inserts a User into the database
// func (c *Client) SaveUser(u *account.User) error {
// 	transactItems := []*dynamodb.TransactWriteItem{}
// 	for _, k := range u.APIKeys {
// 		item, err := dynamodbattribute.MarshalMap(&APIKey{
// 			itemKey: itemKey{
// 				PK: u.ID.String(),
// 				SK: fmt.Sprintf("%s#%s", skPrefixAPIKey, k.Key),
// 			},
// 			Key:      k.Key,
// 			ReadOnly: k.ReadOnly,
// 		})
// 		if err != nil {
// 			return fmt.Errorf("could not marshal APIKey to DynamoDB item")
// 		}

// 		transactItems = append(transactItems, &dynamodb.TransactWriteItem{
// 			Put: &dynamodb.Put{
// 				TableName: aws.String(endpointTableName),
// 				Item:      item,
// 			},
// 		})
// 	}

// 	user, err := dynamodbattribute.MarshalMap(&User{
// 		itemKey: itemKey{
// 			PK: u.ID.String(),
// 			SK: skPrefixUser,
// 		},
// 	})
// 	if err != nil {
// 		return fmt.Errorf("could not marshal user info to User item")
// 	}
// 	transactItems = append(transactItems, &dynamodb.TransactWriteItem{
// 		Put: &dynamodb.Put{
// 			TableName:           aws.String(endpointTableName),
// 			Item:                user,
// 			ConditionExpression: aws.String(fmt.Sprintf("attribute_not_exists(%s)", endpointSK)),
// 		},
// 	})

// 	_, err = c.db.TransactWriteItems(&dynamodb.TransactWriteItemsInput{
// 		TransactItems: transactItems,
// 	})
// 	if err != nil {
// 		if awserr, ok := err.(awserr.Error); ok && awserr.Code() == dynamodb.ErrCodeTransactionCanceledException {
// 			return ErrorConflict
// 		}
// 		return fmt.Errorf("could not persist user to database: %w", err)
// 	}

// 	return nil
// }

// // GetUser retrieves a User from the database using the given ID
// func (c *Client) GetUser(uid account.ID) (*account.User, error) {
// 	uidKey := ":uid"
// 	res, err := c.db.Query(&dynamodb.QueryInput{
// 		TableName:              aws.String(endpointTableName),
// 		KeyConditionExpression: aws.String(fmt.Sprintf("%s = %s", endpointPK, uidKey)),
// 		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
// 			uidKey: {
// 				S: aws.String(uid.String()),
// 			},
// 		},
// 	})
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to retrieve user from database: %w", err)
// 	}
// 	if len(res.Items) == 0 {
// 		return nil, ErrorItemNotFound
// 	}

// 	u := account.User{}
// 	keys := []*account.APIKey{}
// 	for _, i := range res.Items {
// 		var sk string
// 		dynamodbattribute.Unmarshal(i[endpointSK], &sk)
// 		switch strings.Split(sk, "#")[0] {
// 		case skPrefixUser:
// 			userItem := User{}
// 			dynamodbattribute.UnmarshalMap(i, &userItem)
// 			u.ID = account.ID(userItem.PK)
// 		case skPrefixAPIKey:
// 			key := account.APIKey{}
// 			dynamodbattribute.UnmarshalMap(i, &key)
// 			keys = append(keys, &key)
// 		}
// 	}
// 	u.APIKeys = keys

// 	return &u, nil
// }
