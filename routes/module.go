package routes

import (
	"cybercampus_module/controllers"

	"github.com/gofiber/fiber/v2"
)

func ModuleRoute(app *fiber.App) {
	

	app.Get("/modules", controllers.GetAllModules)
	app.Get("/modules/:id", controllers.GetModuleByID)
	app.Post("/modules/add", controllers.CreateModule)
	app.Put("/modules/:id/update", controllers.UpdateModule)
	app.Delete("/modules/:id/delete", controllers.DeleteModule)

}