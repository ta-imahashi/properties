package model

type Counter struct {
	Table string `dynamodbav:"table" label:"テーブル名"`
	Count string `dynamodbav:"count" label:"increment番号"`
}
