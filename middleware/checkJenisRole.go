package middleware

import (
	//"context"
	//"cybercampus_module/configs"
	//"cybercampus_module/models"
	//"time"

	"github.com/gofiber/fiber/v2"
	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/mongo"
)


func CheckJenisRole(role []string) fiber.Handler{
	return func(c *fiber.Ctx) error {

		
		roleClaims := c.Locals("role").(string)


		isAllowed := false 
		for _, v :=range role {
			if v == roleClaims {
				isAllowed = true
				break
			}

			if !isAllowed {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"status":  fiber.StatusUnauthorized,
					"message": "You are not allowed to access this resource",
					"data":    nil,
				})
			}
		}

		return c.Next()
	}
}