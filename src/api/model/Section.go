package model

type Section struct {
	Id         int    `dynamodbav:"id" json:"id" label:"ID"`
	PropertyId int    `dynamodbav:"property_id" json:"property_id" label:"物件ID"`
	Name       string `dynamodbav:"name" json:"name" label:"区画名"`
	Type       int    `dynamodbav:"type" json:"type" label:"種別"`
	Area       int    `dynamodbav:"area" json:"area" label:"面積(m^2)" discription:"小数点位置はSquareで管理"`
	Square     int    `dynamodbav:"square" json:"square" label:"面積の小数点位置" discription:"square=2,area=111のとき、計算後の値は1.11となる"`
	Floor      *int   `dynamodbav:"floor" json:"floor" label:"階数"`
	FloorPlan  *int   `dynamodbav:"floor_plan" json:"floor_plan" label:"間取り"`
	CreatedAt  string `dynamodbav:"created_at" json:"created_at" label:"作成日"`
	UpdatedAt  string `dynamodbav:"updated_at" json:"updated_at" label:"最終更新日"`
}
