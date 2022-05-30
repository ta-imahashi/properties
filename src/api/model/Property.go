package model

type Property struct {
	Id        int    `dynamodbav:"id" json:"id" label:"ID"`
	Name      string `dynamodbav:"name" json:"name" label:"物件名"`
	CreatedAt string `dynamodbav:"created_at" json:"created_at" label:"作成日"`
	UpdatedAt string `dynamodbav:"updated_at" json:"updated_at" label:"最終更新日"`
}
