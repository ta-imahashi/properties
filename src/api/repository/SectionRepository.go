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

type SectionRepository struct {
	Db        *dynamodb.Client
	TableName string
}

func NewSectionRepository(dynamo db.Db) SectionRepository {
	return SectionRepository{Db: dynamo.Client, TableName: "sections"}
}

func (re *SectionRepository) Scan() (items []model.Section) {
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

func (re *SectionRepository) Query(propertyId string) (items []model.Section) {
	index := "property_id_index"
	keyConditionExpression := "property_id = :property_id"

	output, err := re.Db.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              &re.TableName,
		IndexName:              &index,
		KeyConditionExpression: &keyConditionExpression,
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":property_id": &types.AttributeValueMemberN{Value: propertyId},
		},
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

func (re *SectionRepository) GetItem(id string) (item model.Section) {
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

func (re *SectionRepository) PutItem(section model.Section) model.Section {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(jst).Format(time.RFC3339)

	section.CreatedAt = now
	section.UpdatedAt = now

	data, err := attributevalue.MarshalMap(section)

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

	return section
}

func (re *SectionRepository) UpdateItem(section model.Section) (item model.Section) {
	jst, _ := time.LoadLocation("Asia/Tokyo")

	updateExpression := "set #property_id = :property_id, #type = :type,  #name = :name,  #updated_at = :updated_at"

	output, err := re.Db.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: &re.TableName,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberN{Value: strconv.Itoa(section.Id)},
		},
		UpdateExpression: &updateExpression,
		ExpressionAttributeNames: map[string]string{
			"#property_id": "property_id",
			"#type":        "type",
			"#name":        "name",
			"#updated_at":  "updated_at",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":property_id": &types.AttributeValueMemberN{Value: strconv.Itoa(section.PropertyId)},
			":type":        &types.AttributeValueMemberN{Value: strconv.Itoa(section.Type)},
			":name":        &types.AttributeValueMemberS{Value: section.Name},
			":updated_at":  &types.AttributeValueMemberS{Value: time.Now().In(jst).Format(time.RFC3339)},
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

func (re *SectionRepository) DeleteItem(id string) {
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
