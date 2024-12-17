package routes

import (
	"cybercampus_module/controllers"
	"cybercampus_module/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserModuleRoute(app *fiber.App){
	apiUserModule := app.Group("/api", middleware.JwtMiddleware, middleware.CheckJenisRole([]string{"admin"}))
	apiUserModule.Get("/user_moduleAll", controllers.UserModuleFindAll)
	apiUserModule.Post("/user_module/add", controllers.UserModuleAddModule)
	apiUserModule.Delete("/user_module/delete", controllers.UserModuleDeleteModule)

}