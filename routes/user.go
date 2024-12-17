package routes

import (
	"cybercampus_module/controllers"
	"cybercampus_module/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App){

	
	// User access
	app.Post("/login", controllers.Login)
	app.Get("/usersbyId", controllers.GetUserById, middleware.JwtMiddleware)
	app.Get("/user/module", controllers.UserModuleFindByUser, middleware.JwtMiddleware)

	app.Post("/add", controllers.CreateUser)
	verifyTokenAdminGroup := app.Group("/users", middleware.JwtMiddleware, middleware.CheckJenisRole([]string{"admin"}))
	verifyTokenAdminGroup.Get("/" ,controllers.GetAllUsers)
	//verifyTokenAdminGroup.Post("/add", controllers.CreateUser)
	verifyTokenAdminGroup.Put("/:id/update", controllers.UpdateUser)
	verifyTokenAdminGroup.Delete("/:id/delete", controllers.DeleteUser)
	verifyTokenAdminGroup.Post("/changeJenisUser", controllers.UpdateJenisUser)
}