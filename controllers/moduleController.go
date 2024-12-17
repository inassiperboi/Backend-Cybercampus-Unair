package controllers

import (
	"context"
	"cybercampus_module/configs"
	"cybercampus_module/models"
	"cybercampus_module/response"
	"reflect"
	"time"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collectionModule *mongo.Collection = configs.GetCOllection(configs.Client, "module")


func GetAllModules(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var modules []models.ModuleResponse

	cursor, err := collectionModule.Find(ctx, bson.D{})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
			Status:  fiber.StatusInternalServerError,
			Message: "Error when fetching modules",
			Data:    err.Error(),
		})
	}

	if err = cursor.All(ctx, &modules); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
			Status:  fiber.StatusInternalServerError,
			Message: "Error when decoding modules",
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    modules,
	})
}


func GetModuleByID(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	hexId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
			Status:  fiber.StatusInternalServerError,
			Message: "Error when parsing ID",
			Data:    err.Error(),
		})
	}

	var module models.ModuleResponse

	err = collectionModule.FindOne(ctx, bson.M{"_id": hexId}).Decode(&module)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
			Status:  fiber.StatusInternalServerError,
			Message: "Error when fetching module",
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    module,
	})

}

func CreateModule(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	var module models.ModuleRequest

	if err:= c.BodyParser(&module); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status:  fiber.StatusBadRequest,
			Message: "Error when parsing request",
			Data:    err.Error(),
		})
	}

	var lastModule models.ModuleResponse

	err := collectionModule.FindOne(ctx, bson.M{}, options.FindOne().SetSort(bson.M{"urutan": -1})).Decode(&lastModule)

	if err != nil || lastModule.URUTAN == 0 {
		module.URUTAN = 1
	} else {
		module.URUTAN = lastModule.URUTAN + 1

	}

	newModule := models.ModuleRequest{
		NAMA_MODULE: module.NAMA_MODULE,
		KETERANGAN: module.KETERANGAN,
		URUTAN:     module.URUTAN,
		ICON:       module.ICON,
		STATUS:     true,
		CREATED_AT: time.Now().Format("2006-01-02 15:04:05"),
		UPDATED_AT: time.Now().Format("2006-01-02 15:04:05"),
	}

	result, err := collectionModule.InsertOne(ctx, newModule)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
			Status:  fiber.StatusInternalServerError,
			Message: "Error when inserting module",
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(response.Response{
		Status: fiber.StatusCreated,
		Message: "Success",
		Data:    result,
	})

}


func UpdateModule(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	hexId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
			Status:  fiber.StatusInternalServerError,
			Message: "Error when parsing ID",
			Data:    err.Error(),
		})
	}

	var module models.ModuleRequest

	if err := c.BodyParser(&module); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status:  fiber.StatusBadRequest,
			Message: "Error when parsing request",
			Data:    err.Error(),
		})
	}

	updateModule := bson.M{}

	value := reflect.ValueOf(module)
	tipe := reflect.TypeOf(module)
	
	for i := 0; i < value.NumField(); i++{

		fieldName := tipe.Field(i).Tag.Get("bson")

		if fieldName == "" || fieldName == "-" {
			continue
		}

		fieldValue := value.Field(i).Interface()

		switch val := fieldValue.(type){
		case string:
			if val != ""{
				updateModule[fieldName] = val
			}
		case bool:
			if val {
				updateModule[fieldName] = val
			}
		case time.Time:
			if !val.IsZero(){
				updateModule[fieldName] = val
			}
		default:
			if !reflect.ValueOf(fieldValue).IsZero(){
				updateModule[fieldName] = fieldValue
			}
		}


	}

	if len(updateModule) > 0 {
		updateModule["updated_at"] = time.Now().Format("2006-01-02 15:04:05")
		result, err := collectionModule.UpdateOne(ctx, bson.M{"_id": hexId}, bson.M{"$set": updateModule})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
				Status:  fiber.StatusInternalServerError,
				Message: "Error when updating module",
				Data:    err.Error(),
			})
		}

		if result.MatchedCount == 0 {
			return c.Status(fiber.StatusNotFound).JSON(response.Response{
				Status:  fiber.StatusNotFound,
				Message: "Module not found",
				Data:    nil,
			})
		}
	}else{
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status: fiber.StatusBadRequest,
			Message: "No field to update",
			Data: nil,
		})
	}


	return c.Status(fiber.StatusOK).JSON(response.Response{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    nil,
	})
}

func DeleteModule(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	hex_Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
			Status:  fiber.StatusInternalServerError,
			Message: "Error when parsing ID",
			Data:    err.Error(),
		})
	}

	result, err := collectionModule.DeleteOne(ctx, bson.M{"_id": hex_Id})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
			Status:  fiber.StatusInternalServerError,
			Message: "Error when deleting module",
			Data:    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}