package helpers

import (
	"context"
	"cybercampus_module/configs"
	"cybercampus_module/models"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionUserModule = configs.GetCOllection(configs.Client, "user_module")
var collectionTemplate = configs.GetCOllection(configs.Client, "templates")


func SyncModuleTemplate(jenis_user primitive.ObjectID, idUser primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("Syncing module for user with id:", idUser)
	fmt.Println("Jenis user:", jenis_user)

	
	var userModule models.UserModule
	err := collectionUserModule.FindOne(ctx, bson.M{"id_user": idUser}).Decode(&userModule)

	if err != nil {
	
		if err.Error() == "mongo: no documents in result" {
			
			var template models.TemplateUserModuleRequest
			err = collectionTemplate.FindOne(ctx, bson.M{"_id": jenis_user}).Decode(&template)

			fmt.Println("Template:", template)

			if err != nil {
				return false, err
			}

			
			newModules := []primitive.ObjectID{}
			newModules = append(newModules, template.Template...)
	
			newUserModule := models.UserModule{
				ID:         primitive.NewObjectID(),
				IDUser:     idUser,
				JenisUser:  template.JenisUser,
				MODULES:    newModules,
				CREATED_AT: time.Now(),
				UPDATED_AT: time.Now(),
			}

			_, err = collectionUserModule.InsertOne(ctx, newUserModule)
			if err != nil {
				return false, err
			}

			fmt.Println("User module created : ", newUserModule)
			return true, nil
		}

		
		return false, err
	}

	
	var template models.TemplateUserModuleRequest
	err = collectionTemplate.FindOne(ctx, bson.M{"_id": jenis_user}).Decode(&template)

	if err != nil {
		return false, err
	}


	updatedModules := []primitive.ObjectID{}
	updatedModules = append(updatedModules, template.Template...)
	
	update := bson.M{
		"$set": bson.M{
			"modules":    updatedModules,
			"jenis_user": template.JenisUser,
			"updated_at": time.Now(),
		},
	}

	_, err = collectionUserModule.UpdateOne(ctx, bson.M{"id_user": idUser}, update)
	if err != nil {
		return false, err
	}

	return true, nil
}


func SyncUpdateTemplate(jenis_user string, template []primitive.ObjectID) (bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var checkJenisUserModule models.UserModule
	err := collectionTemplate.FindOne(ctx, bson.M{"jenis_user": jenis_user}).Decode(&checkJenisUserModule)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	update := bson.M{
		"$set" : bson.M{
			"modules" : template,
		},
	}

	filter := bson.M{
		"jenis_user" : jenis_user,
	}

	_, err = collectionUserModule.UpdateMany(ctx, filter, update)

	if err != nil {
		return false, err
	}

	return true, nil
	
}