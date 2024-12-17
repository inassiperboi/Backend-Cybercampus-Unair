package controllers

import (
	"context"
	"cybercampus_module/configs"
	"cybercampus_module/models"

	//	"fmt"

	//"cybercampus_module/models"
	"cybercampus_module/response"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionUserModule *mongo.Collection = configs.GetCOllection(configs.Client, "user_module")

func UserModuleFindAll(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	pipeline := mongo.Pipeline{
		{ {
			Key : "$lookup", Value : bson.D{
				{Key : "from", Value : "users"},
				{Key : "localField", Value : "id_user"},
				{Key : "foreignField", Value : "_id"},
				{Key : "as", Value : "user_data"},
			},
		},},
		{{ 
			Key : "$unwind", Value : "$user_data",
		 }},
		 {{ 
			Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "module"},
				{Key: "localField", Value: "modules"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "module_data"},
			},
		  }},
		  {{ 
			Key: "$project", Value: bson.D{
				{Key: "_id", Value: 1},
				{Key: "id_user", Value: 1},
				{Key: "jenis_user", Value: 1},
				{Key: "created_at", Value: 1},
				{Key: "updated_at", Value: 1},
				{Key: "username", Value: "$user_data.username"},
				{Key: "nm_user", Value: "$user_data.nm_user"},
				{Key: "email", Value: "$user_data.email"},
				{Key: "password", Value: "$user_data.password"},
				{Key: "role", Value: "$user_data.role"},
				{Key: "modules", Value: "$module_data"},
			},
		   }},
	}

	cursor, err := collectionUserModule.Aggregate(ctx, pipeline)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Error when fetching aggregrate user module",
			"data":    err.Error(),
		})
	}

	

	var userModuleResult []models.UserModuleResponse

	if err = cursor.All(ctx, &userModuleResult); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Error when decoding user module",
			"data":    err.Error(),
		})
	}

	

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    userModuleResult,	
	})
}


func UserModuleFindByUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()


	idUser := c.Locals("id_user").(string)
	role := c.Locals("role").(string)

	if role == "admin" {
		
		cursor, err := collectionModule.Find(ctx, bson.M{})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  fiber.StatusInternalServerError,
				"message": "Error when fetching modules",
				"data":    err.Error(),
			})
		}

		var modules []models.ModuleResponse
		if err = cursor.All(ctx, &modules); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  fiber.StatusInternalServerError,
				"message": "Error when decoding modules",
				"data":    err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(response.Response{
			Status:  fiber.StatusOK,
			Message: "Success",
			Data:    modules,
		})
	}

	hexID, _ := primitive.ObjectIDFromHex(idUser)
	if idUser == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "id_user is required",
			"data":    nil,
		})
	}

	
	pipeline := mongo.Pipeline{
		{
			{Key: "$match", Value: bson.D{
				{Key: "id_user", Value: hexID},
			}},
		},
		{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "users"},
				{Key: "localField", Value: "id_user"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "user_data"},
			}},
		},
		{
			{Key: "$unwind", Value: "$user_data"},
		},
		{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "module"},
				{Key: "localField", Value: "modules"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "module_data"},
			}},
		},
		{
			{Key: "$project", Value: bson.D{
				{Key: "_id", Value: 1},
				{Key: "id_user", Value: 1},
				{Key: "jenis_user", Value: 1},
				{Key: "created_at", Value: 1},
				{Key: "updated_at", Value: 1},
				{Key: "username", Value: "$user_data.username"},
				{Key: "nm_user", Value: "$user_data.nm_user"},
				{Key: "email", Value: "$user_data.email"},
				{Key: "password", Value: "$user_data.password"},
				{Key: "role", Value: "$user_data.role"},
				{Key: "modules", Value: "$module_data"},
			}},
		},
	}

	
	cursor, err := collectionUserModule.Aggregate(ctx, pipeline)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Error when fetching aggregate user module",
			"data":    err.Error(),
		})
	}


	var userModuleResult []models.UserModuleResponse
	if err = cursor.All(ctx, &userModuleResult); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Error when decoding user module",
			"data":    err.Error(),
		})
	}


	if len(userModuleResult) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "No user module found for the given id_user",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    userModuleResult,
	})
}



func UserModuleAddModule (c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var userModule models.UserModule

	if err := c.BodyParser(&userModule); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
			Status: fiber.StatusInternalServerError,
			Message: "Error when parsing body",
			Data: err.Error(),
		})
	}

	if userModule.MODULES == nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status: fiber.StatusBadRequest,
			Message: "Error Module Not Found",
			Data: "MODULES is required",
		})
	}

	updateuserModule := bson.M{
		"$push" :bson.M{
			"modules" : bson.M{"$each": userModule.MODULES},
		},
	}

	filter := bson.M{"id_user" : userModule.IDUser}

	result , err := collectionUserModule.UpdateOne(ctx, filter, updateuserModule)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
			Status: fiber.StatusInternalServerError,
			Message: "Error when updating user module",
			Data: err.Error(),
		})
	}

	if result.ModifiedCount == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status: fiber.StatusBadRequest,
			Message: "Error when updating user module",
			Data: "User module not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Status: fiber.StatusOK,
		Message: "Success",
		Data: result,
	})
}


func UserModuleDeleteModule (c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var userModule models.UserModule

	if err := c.BodyParser(&userModule); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
			Status: fiber.StatusInternalServerError,
			Message: "Error when parsing body",
			Data: err.Error(),
		})
	}

	if userModule.MODULES == nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status: fiber.StatusBadRequest,
			Message: "Error Module Not Found",
			Data: "MODULES is required",
		})
	}

	updateuserModule := bson.M{
		"$pull" :bson.M{
			"modules": bson.M{"$in": userModule.MODULES},
		},
	}


	filter := bson.M{"id_user" : userModule.IDUser}

	result , err := collectionUserModule.UpdateOne(ctx, filter, updateuserModule)


	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
			Status: fiber.StatusInternalServerError,
			Message: "Error when updating user module",
			Data: err.Error(),
		})
	}

	if result.ModifiedCount == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status: fiber.StatusBadRequest,
			Message: "Error when updating user module",
			Data: "User module not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Status: fiber.StatusOK,
		Message: "Success",
		Data: result,
	})
}