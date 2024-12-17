package routes

import (
	"cybercampus_module/controllers"
	//"cybercampus_module/middleware"

	"github.com/gofiber/fiber/v2"
)

func JenisRoleRoute(app *fiber.App) {
	
	// apiJenisUser := app.Group("/api")
	app.Get("/jenis_roles", controllers.GetAllJenisUser)
	app.Post("/jenis_roles/add", controllers.CreateJenisUser)
}