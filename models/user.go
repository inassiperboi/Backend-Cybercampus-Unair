package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type UserResponse struct {
	ID                  primitive.ObjectID    `bson:"_id,omitempty" json:"ID"`
	Username            string    `bson:"username" json:"USERNAME"`
	NM_USER            	string    `bson:"nm_user" json:"NM_USER"`
	Password            string    `bson:"password" json:"PASSWORD"`
	Email               string    `bson:"email" json:"EMAIL"`
	Role            	string    `bson:"role" json:"ROLE"`
	CreatedAt           time.Time `bson:"created_at" json:"CREATED_AT"`
	UpdatedAt           time.Time `bson:"updated_at" json:"UPDATED_AT"`
	IsActive            bool      `bson:"is_active" json:"IS_ACTIVE"`
	LastLogin           time.Time `bson:"last_login,omitempty" json:"LAST_LOGIN,omitempty"`
	Photo               string    `bson:"photo,omitempty" json:"PHOTO,omitempty"`
	Phone               string    `bson:"phone,omitempty" json:"PHONE,omitempty"`
	JENIS_USER          primitive.ObjectID    `bson:"jenis_user,omitempty" json:"JENIS_USER,omitempty"`
	Address             string    `bson:"address,omitempty" json:"ADDRESS,omitempty"`
	Gender              string    `bson:"gender,omitempty" json:"GENDER,omitempty"` 
	DateOfBirth         time.Time `bson:"tanggal_lahir,omitempty" json:"TANGGAL_LAHIR,omitempty"`

}

type UserRequest struct {
	ID                  primitive.ObjectID    `bson:"_id,omitempty" json:"ID"`
	Username            string    `bson:"username" json:"USERNAME"`
	NM_USER            	string    `bson:"nm_user" json:"NM_USER"`
	Password            string    `bson:"password" json:"PASSWORD"`
	Email               string    `bson:"email" json:"EMAIL"`
	Role            	string    `bson:"role" json:"ROLE"`
	CreatedAt           time.Time `bson:"created_at" json:"CREATED_AT"`
	UpdatedAt           time.Time `bson:"updated_at" json:"UPDATED_AT"`
	IsActive            bool      `bson:"is_active" json:"IS_ACTIVE"`
	LastLogin           time.Time `bson:"last_login" json:"LAST_LOGIN"`
	Photo               string    `bson:"photo" json:"PHOTO"`
	Phone               string    `bson:"phone" json:"PHONE"`
	JENIS_USER          primitive.ObjectID    `bson:"jenis_user" json:"JENIS_USER"`
	Address             string    `bson:"address" json:"ADDRESS"`
	Gender              string    `bson:"gender" json:"GENDER"`
	DateOfBirth         time.Time `bson:"tanggal_lahir" json:"TANGGAL_LAHIR"`

}


type UserLogin struct {
	ID                  string    `bson:"_id,omitempty" json:"ID"`
	Username            string    `bson:"username" json:"USERNAME"`
	NM_USER            	string    `bson:"nm_user" json:"NM_USER"`
	Password            string    `bson:"password" json:"PASSWORD"`
	Email               string    `bson:"email" json:"EMAIL"`
	Role            	string    `bson:"role" json:"ROLE"`
	CreatedAt           time.Time `bson:"created_at" json:"CREATED_AT"`
	UpdatedAt           time.Time `bson:"updated_at" json:"UPDATED_AT"`
	IsActive            bool      `bson:"is_active" json:"IS_ACTIVE"`
	LastLogin           time.Time `bson:"last_login,omitempty" json:"LAST_LOGIN,omitempty"`
	Photo               string    `bson:"photo,omitempty" json:"PHOTO,omitempty"`
	Phone               string    `bson:"phone,omitempty" json:"PHONE,omitempty"`
	JENIS_USER          string    `bson:"jenis_user,omitempty" json:"JENIS_USER,omitempty"`
	Address             string    `bson:"address,omitempty" json:"ADDRESS,omitempty"`
	Gender              string    `bson:"gender,omitempty" json:"GENDER,omitempty"` 
	DateOfBirth         time.Time `bson:"tanggal_lahir,omitempty" json:"TANGGAL_LAHIR,omitempty"`
	TOKEN 			    string    `bson:"token,omitempty" json:"TOKEN,omitempty"` 
}