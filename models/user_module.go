package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModule struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	IDUser     primitive.ObjectID `json:"id_user" bson:"id_user"`
	JenisUser  string 			`json:"jenis_user" bson:"jenis_user"`
	MODULES    []primitive.ObjectID   `json:"modules" bson:"modules"`
	CREATED_AT time.Time             `json:"created_at" bson:"created_at"`
	UPDATED_AT time.Time `json:"updated_at" bson:"updated_at"`
}

type UserModuleResponse struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	IDUser     primitive.ObjectID `json:"id_user" bson:"id_user"`
	Username   string             `bson:"username" json:"USERNAME"`
	NM_USER    string             `bson:"nm_user" json:"NM_USER"`
	Email      string             `bson:"email" json:"EMAIL"`
	Password   string             `bson:"password" json:"PASSWORD"`
	Role       string             `bson:"role" json:"ROLE"`
	JenisUser           string   `json:"jenis_user" bson:"jenis_user"`
	Module              []ModuleResponse `json:"modules" bson:"modules"`
	Created_at time.Time `json:"created_at" bson:"created_at"`
	Updated_at time.Time `json:"updated_at" bson:"updated_at"`
}