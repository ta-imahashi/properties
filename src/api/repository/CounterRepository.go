package repository

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ta-imahashi/properties/db"
	"github.com/ta-imahashi/properties/model"
)

type CounterRepository struct {
	Db        *dynamodb.Client
	TableName string
}

func NewCounterRepository(dynamo db.Db) CounterRepository {
	return CounterRepository{Db: dynamo.Client, TableName: "counters"}
}

func (re *CounterRepository) Increment(tableName string) (counter model.Counter) {
	output, err := re.Db.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(re.TableName),
		Key: map[string]types.AttributeValue{
			"table": &types.AttributeValueMemberS{Value: tableName},
		},
		UpdateExpression: aws.String("ADD #count :incr"),
		ExpressionAttributeNames: map[string]string{
			"#count": "count",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":incr": &types.AttributeValueMemberN{Value: "1"},
		},
		ReturnValues: types.ReturnValueUpdatedNew,
	})

	if err != nil {
		log.Fatalf("Counter: %#v", err)
	}

	err = attributevalue.UnmarshalMap(output.Attributes, &counter)

	if err != nil {
		log.Fatalf("Unmarshal: %#v", err)
	}

	return
}
