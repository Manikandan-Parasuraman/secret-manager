package models

type Secret struct {
	ID     string `json:"id" bson:"_id,omitempty"`
	Name   string `json:"name" bson:"name"`
	Secret string `json:"secret" bson:"secret"`
}
