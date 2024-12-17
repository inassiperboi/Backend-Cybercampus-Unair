package middleware

import (
	"cybercampus_module/configs"
	"cybercampus_module/helpers"
	"cybercampus_module/models"
	"cybercampus_module/response"
	"encoding/base64"
	"encoding/json"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func JwtMiddleware(c *fiber.Ctx) error {
	secretKey := os.Getenv(configs.LoadEnv("SECRET_KEY"))

	token := c.Get("Authorization")

	if strings.HasPrefix(token, "Bearer"){
		token = strings.TrimPrefix(token, "Bearer ")
	}else {
		return c.Status(fiber.StatusUnauthorized).JSON(response.Response{
			Status: fiber.StatusUnauthorized,
			Message: "Unauthorized",
			Data: "Unauthorized",
		})
	}

	parts := strings.Split(token, ".")

	if len(parts) != 3 {
		return c.Status(fiber.StatusUnauthorized).JSON(response.Response{
			Status: fiber.StatusUnauthorized,
			Message: "Invalid Token !",
			Data: "Invalid Token !",
		})
	}

	_, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.Response{
			Status: fiber.StatusUnauthorized,
			Message: "Invalid Header Token !",
			Data: "Invalid Header Token !",
	})
}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.Response{
			Status: fiber.StatusUnauthorized,
			Message: "Invalid Payload Token !",
			Data: "Invalid Payload Token !",
		})
	}

	_ , err = base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Response{
			Status: fiber.StatusInternalServerError,
			Message: "Invalid Signature",
			Data: nil,
		})
	}

	expectedSignature := helpers.CreateSignature(parts[0], parts[1], secretKey)
	if expectedSignature != parts[2] {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "Invalid token signature",
		})
	}

	var claims models.JWTClaims

	if err := json.Unmarshal(payload, &claims); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Error when parsing token",
			"data":    err.Error(),
		})
	}

	if time.Unix(claims.Exp,0).Before(time.Now()){
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "Token has expired",
		})
	}
	
	c.Locals("id_user", claims.ID)
	c.Locals("username", claims.Username)
	c.Locals("email", claims.Email)
	c.Locals("jenis_user", claims.JenisUser)
	c.Locals("role", claims.Role)
	return c.Next()
}