package routes

import (
	"cybercampus_module/controllers"
	

	"github.com/gofiber/fiber/v2"
)

func TemplateRoute(app *fiber.App){
	app.Get("/templates", controllers.GetTemplateALL)
	app.Post("/templates/add", controllers.CreateTemplate)
	app.Put("/templates/:id/update", controllers.UpdateTemplate)
}