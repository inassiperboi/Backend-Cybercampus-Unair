package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type ModuleRequest struct {
	ID 	primitive.ObjectID `bson:"_id,omitempty" json:"ID"`
	NAMA_MODULE string `bson:"nama_module" json:"NAMA_MODULE"`
	KETERANGAN string `bson:"keterangan" json:"KETERANGAN"`
	URUTAN int `bson:"urutan" json:"URUTAN"`
	ICON string `bson:"icon" json:"ICON"`
	STATUS bool `bson:"status" json:"STATUS"`
	CREATED_AT string `bson:"created_at" json:"CREATED_AT"`
	UPDATED_AT string `bson:"updated_at" json:"UPDATED_AT"`
}

type ModuleResponse struct {
	ID 	primitive.ObjectID `bson:"_id,omitempty" json:"ID"`
	NAMA_MODULE string `bson:"nama_module" json:"NAMA_MODULE"`
	KETERANGAN string `bson:"keterangan" json:"KETERANGAN"`
	URUTAN int `bson:"urutan" json:"URUTAN"`
	ICON string `bson:"icon, omitempty" json:"ICON"`
	STATUS bool `bson:"status" json:"STATUS"`
	CREATED_AT string `bson:"created_at" json:"CREATED_AT"`
	UPDATED_AT string `bson:"updated_at" json:"UPDATED_AT"`
}