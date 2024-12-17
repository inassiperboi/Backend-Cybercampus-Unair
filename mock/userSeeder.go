package mock

import (
	"context"
	"cybercampus_module/configs"
	"cybercampus_module/helpers"
	"cybercampus_module/models"
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection = configs.GetCOllection(configs.Client, "users")

func UserSeeder() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()


	HashPasword := helpers.HashPasword("admin555")

	newUser := models.UserRequest{
		Username: "admin",
		NM_USER: "343249385",
		Password: HashPasword,
		Email:    "admin@gmail.com",
		Role:    "admin",
		IsActive: true,
	}

	_, err := collection.InsertOne(ctx, newUser)

	if err != nil {
		panic(err)
	}

	fmt.Print("User Seeder Admin Created")
	
}