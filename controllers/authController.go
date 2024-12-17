package controllers

import (
	"context"
	"cybercampus_module/helpers"
	"cybercampus_module/models"
	"cybercampus_module/response"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func Login(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.UserRequest

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid request",
			"data":    err.Error(),
		})
	}

	var LoginData models.UserRequest

	err := collection.FindOne(ctx, bson.M{"email" : user.Email}).Decode(&LoginData)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "User Not Found",
			"data":    err.Error(),
		})
	}

	if !helpers.ComparePassword(LoginData.Password, user.Password) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Password is incorrect",
			"data":    nil,
		})
	}
	
	var jenisUser models.JenisUserResponse

	_ = collectionTemplate.FindOne(ctx, bson.M{"_id": LoginData.JENIS_USER}).Decode(&jenisUser)

	token, err  := helpers.GenerateToken(LoginData.ID.Hex(), LoginData.Username, LoginData.Email, jenisUser.JenisUser, LoginData.Role)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Error when generating token",
			"data":    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    models.UserLogin{
			ID:        LoginData.ID.Hex(),
			Username:  LoginData.Username,
			NM_USER:  LoginData.NM_USER,
			Email:     LoginData.Email,
			Password: LoginData.Password,
			JENIS_USER: func() string {
				if jenisUser.JenisUser == "" {
					return "admin"
				}
				return jenisUser.JenisUser
			}(),
			IsActive: LoginData.IsActive,
			Role: 	LoginData.Role,
			Phone: 	LoginData.Phone,
			Address: 	LoginData.Address,
			DateOfBirth: LoginData.DateOfBirth,
			TOKEN:     token,
		} , 
	},
	) 
	
}

