package repository

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ta-imahashi/properties/db"
	"github.com/ta-imahashi/properties/model"
)

type PropertyRepository struct {
	Db        *dynamodb.Client
	TableName string
}

func NewPropertyRepository(dynamo db.Db) PropertyRepository {
	return PropertyRepository{Db: dynamo.Client, TableName: "properties"}
}

func (re *PropertyRepository) Scan() (items []model.Property) {
	output, err := re.Db.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: &re.TableName,
	})

	if err != nil {
		log.Fatalf("%#v", err)
	}

	err = attributevalue.UnmarshalListOfMaps(output.Items, &items)

	if err != nil {
		log.Fatalf("%#v", err)
	}

	return items
}

func (re *PropertyRepository) GetItem(id string) (item model.Property) {
	output, err := re.Db.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: &re.TableName,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberN{Value: id},
		},
	})

	if err != nil {
		log.Fatalf("%#v", err)
	}

	err = attributevalue.UnmarshalMap(output.Item, &item)

	if err != nil {
		log.Fatalf("%#v", err)
	}

	return item
}

func (re *PropertyRepository) PutItem(property model.Property) model.Property {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(jst).Format(time.RFC3339)

	property.CreatedAt = now
	property.UpdatedAt = now

	data, err := attributevalue.MarshalMap(property)

	if err != nil {
		log.Fatalf("Marshal: %#v", err)
	}

	_, err = re.Db.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: &re.TableName,
		Item:      data,
	})

	if err != nil {
		log.Fatalf("%#v", err)
	}

	return property
}

func (re *PropertyRepository) UpdateItem(property model.Property) (item model.Property) {
	jst, _ := time.LoadLocation("Asia/Tokyo")

	updateExpression := "set #name = :name, #create_at = :created_at"

	output, err := re.Db.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: &re.TableName,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberN{Value: strconv.Itoa(property.Id)},
		},
		UpdateExpression: &updateExpression,
		ExpressionAttributeNames: map[string]string{
			"#name":       "name",
			"#created_at": "created_at",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":name":       &types.AttributeValueMemberS{Value: property.Name},
			":created_at": &types.AttributeValueMemberS{Value: time.Now().In(jst).Format(time.RFC3339)},
		},
		ReturnValues: types.ReturnValueAllNew,
	})

	if err != nil {
		log.Fatalf("%#v", err)
	}

	err = attributevalue.UnmarshalMap(output.Attributes, &item)

	if err != nil {
		log.Fatalf("%#v", err)
	}

	return item
}

func (re *PropertyRepository) DeleteItem(id string) {
	_, err := re.Db.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: &re.TableName,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberN{Value: id},
		},
	})

	if err != nil {
		log.Fatalf("%#v", err)
	}
}
