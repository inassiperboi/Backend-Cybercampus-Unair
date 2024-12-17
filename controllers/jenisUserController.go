package controllers

import (
	"context"
	"cybercampus_module/models"
	"cybercampus_module/response"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// JENIS ROLE
func GetAllJenisUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var jenisUser []models.JenisUserResponse


	cursor, err := collectionTemplate.Find(ctx, bson.D{})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
			Status: fiber.StatusInternalServerError,
			Message: "Error when fetching jenis user",
			Data: err.Error(),
		})
	}

	if err = cursor.All(ctx, &jenisUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
			Status: fiber.StatusInternalServerError,
			Message: "Error when decoding jenis user",
			Data: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Status: fiber.StatusOK,
		Message: "Success",
		Data: jenisUser,
	})
}

func CreateJenisUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var jenisUser models.TemplateRequest

	if err := c.BodyParser(&jenisUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid request",
			"data":    err.Error(),
		})
	}

	fmt.Print(jenisUser)
	if jenisUser.JenisUser == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Jenis User is required",
			"data":    nil,
		})
	}

	newJenisUser := models.TemplateRequest{
		JenisUser: jenisUser.JenisUser,
		Template: []primitive.ObjectID{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := collectionTemplate.InsertOne(ctx, newJenisUser)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Error when inserting jenis user",
			"data":    err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(response.Response{
		Status: fiber.StatusCreated,
		Message: "Success",
		Data: result,
	})

}

