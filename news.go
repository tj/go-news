// Package news provides a very simple DynamoDB-backed mailing list for newsletters.
package news

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// item model.
type item struct {
	Newsletter string    `json:"newsletter"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"created_at"`
}

// New returns a new mailing list store with default AWS credentials.
func New(table string) *Store {
	return &Store{
		Client:    dynamodb.New(session.New(aws.NewConfig())),
		TableName: table,
	}
}

// Store is a DynamoDB mailing list storage implementation.
type Store struct {
	TableName string
	Client    dynamodbiface.DynamoDBAPI
}

// AddSubscriber adds a subscriber to a newsletter.
func (s *Store) AddSubscriber(newsletter, email string) error {
	i, err := dynamodbattribute.MarshalMap(item{
		Newsletter: newsletter,
		Email:      email,
		CreatedAt:  time.Now(),
	})

	if err != nil {
		return err
	}

	_, err = s.Client.PutItem(&dynamodb.PutItemInput{
		TableName: &s.TableName,
		Item:      i,
	})

	if err != nil {
		return err
	}

	return nil
}

// RemoveSubscriber removes a subscriber from a newsletter.
func (s *Store) RemoveSubscriber(newsletter, email string) error {
	_, err := s.Client.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: &s.TableName,
		Key: map[string]*dynamodb.AttributeValue{
			"newsletter": &dynamodb.AttributeValue{
				S: &newsletter,
			},
			"email": &dynamodb.AttributeValue{
				S: &email,
			},
		},
	})

	return err
}

// GetSubscribers returns subscriber emails for a newsletter.
func (s *Store) GetSubscribers(newsletter string) (emails []string, err error) {
	query := &dynamodb.QueryInput{
		TableName:              &s.TableName,
		AttributesToGet:        aws.StringSlice([]string{"email"}),
		KeyConditionExpression: aws.String(`newsletter = :newsletter`),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":newsletter": &dynamodb.AttributeValue{
				S: &newsletter,
			},
		},
	}

	err = s.Client.QueryPages(query, func(page *dynamodb.QueryOutput, more bool) bool {
		for _, item := range page.Items {
			if v, ok := item["email"]; ok {
				emails = append(emails, *v.S)
			}
		}
		return true
	})

	return
}
