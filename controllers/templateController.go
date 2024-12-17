package controllers

import (
	"context"
	"cybercampus_module/configs"
	"cybercampus_module/helpers"
	"cybercampus_module/models"
	"cybercampus_module/response"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

var collectionTemplate *mongo.Collection = configs.GetCOllection(configs.Client, "templates")

//TEMPLATE SECTION 


func CreateTemplate(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var template models.TemplateRequest
	if err := c.BodyParser(&template); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid request",
			"data":    err.Error(),
		})
	}

	//check jenis user ada tida 
	var checkJenisUser models.TemplateRequest
	err := collectionTemplate.FindOne(ctx, bson.M{"jenis_user": template.JenisUser}).Decode(&checkJenisUser)

	fmt.Print("Check Jenis User",checkJenisUser)

	if err != nil {
		if err == mongo.ErrNoDocuments {
		
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": "Jenis User does not exist",
				"data":    nil,
			})
		}
	
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Error when checking Jenis User",
			"data":    err.Error(),
		})
	}


	if template.JenisUser == ""  {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Jenis User is required",
			"data":    nil,
		})
	}

	if template.Template == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Template is required",
			"data":    nil,
		})
	}

	newTemplate := bson.M{
		"$push": bson.M{
			"template": bson.M{
				"$each": template.Template, 
			},
		},
	}


	result, err := collectionTemplate.UpdateOne(ctx, bson.M{"jenis_user" : template.JenisUser}, newTemplate)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Error when inserting template",
			"data":    err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(response.Response{
		Status: fiber.StatusCreated,
		Message: "Success",
		Data:    result,
	})


}

func GetTemplateALL(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Pipeline untuk agregasi
	pipeline := mongo.Pipeline{
	
		{{
			Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "module"},               
				{Key: "localField", Value: "template"},       
				{Key: "foreignField", Value: "_id"},          
				{Key: "as", Value: "template_details"},       
			},
		}},
	
		{{
			Key: "$project", Value: bson.D{
				{Key: "_id", Value: 1},                      
				{Key: "jenis_user", Value: 1},               
				{Key: "created_at", Value: 1},               
				{Key: "updated_at", Value: 1},               
				{Key: "template", Value: "$template_details"},
			},
		}},
	}

	var rawResults []bson.M

	cursor, err := collectionTemplate.Aggregate(ctx, pipeline)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
			Status:  fiber.StatusInternalServerError,
			Message: "Error when fetching templates",
			Data:    err.Error(),
		})
	}

	fmt.Println("Hasil cursor :", cursor)

	
	if err = cursor.All(ctx, &rawResults); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
			Status:  fiber.StatusInternalServerError,
			Message: "Error when decoding templates",
			Data:    err.Error(),
		})
	}

	fmt.Print("Hasil Raw",rawResults)
	

	var processedResults []models.TemplateResponse
	for _, rawResult := range rawResults {
		
		var templateResponse models.TemplateResponse
		data, _ := bson.Marshal(rawResult)
		bson.Unmarshal(data, &templateResponse)

		
		if templateResponse.Template == nil {
			templateResponse.Template = []models.ModuleResponse{}
		}

		processedResults = append(processedResults, templateResponse)
	}

	// Return response
	return c.Status(fiber.StatusOK).JSON(response.Response{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    processedResults,
	})
}


func UpdateTemplate(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	hexId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
			Status: fiber.StatusInternalServerError,
			Message: "Error when parsing id",
			Data: err.Error(),
		})
	}

	var template models.TemplateRequest

	if err := c.BodyParser(&template); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid request",
			"data":    err.Error(),
		})
	}

	filter := bson.M{"_id" : hexId}

	update := bson.M{
		"$set": bson.M{
			"template": template.Template, 
		},
		}
	

	result, err := collectionTemplate.UpdateOne(ctx, filter, update)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
			Status: fiber.StatusInternalServerError,
			Message: "Error when updating template",
			Data: err.Error(),
		})
	}

	if result.ModifiedCount == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status: fiber.StatusBadRequest,
			Message: "No document was updated",
			Data: nil,
		})
	}

	var jenisRole models.TemplateRequest
	_ = collectionTemplate.FindOne(ctx, bson.M{"_id": hexId}).Decode(&jenisRole)

	// Sync update user module

    cek , err := helpers.SyncUpdateTemplate(jenisRole.JenisUser, jenisRole.Template)

	if err != nil && !cek {
		fmt.Printf("CheckSync Update Template : %t, %s\n", cek, err)
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Status: fiber.StatusOK,
		Message: "Success",
		Data: result,
	})
}

func DeleteTemplate(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	hexId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
			Status: fiber.StatusInternalServerError,
			Message: "Error when parsing id",
			Data: err.Error(),
		})
	}

	result, err := collectionTemplate.DeleteOne(ctx, bson.M{"_id": hexId})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
			Status: fiber.StatusInternalServerError,
			Message: "Error when deleting template",
			Data: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Status: fiber.StatusOK,
		Message: "Success",
		Data: result,
	})
}