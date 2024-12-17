package routes

import (
	"cybercampus_module/controllers"
	"cybercampus_module/middleware"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	// Prefic keseluruhan aja gaes
	api := app.Group("/api")

	// Semua bisa
	api.Post("/login", controllers.Login)

	// User Access 
	user := api.Group("/users", middleware.JwtMiddleware)
	user.Get("/profile", controllers.GetUserById) 
	user.Get("/module", controllers.UserModuleFindByUser) 

	
	admin := api.Group("/admin", middleware.JwtMiddleware, middleware.CheckJenisRole([]string{"admin"}))

	// Admin User Management
	
	adminUsers := admin.Group("/users")
	adminUsers.Get("/", controllers.GetAllUsers)
	adminUsers.Post("/add", controllers.CreateUser)
	adminUsers.Put("/:id/update", controllers.UpdateUser)
	adminUsers.Delete("/:id/delete", controllers.DeleteUser)
	adminUsers.Post("/changeJenisUser", controllers.UpdateJenisUser)

	// Admin Jenis Role Management
	adminJenisRoles := admin.Group("/jenis_roles")
	adminJenisRoles.Get("/", controllers.GetAllJenisUser)
	adminJenisRoles.Post("/add", controllers.CreateJenisUser)

	// Admin Module Management
	adminModules := admin.Group("/modules")
	adminModules.Get("/", controllers.GetAllModules)
	adminModules.Get("/:id", controllers.GetModuleByID)
	adminModules.Post("/add", controllers.CreateModule)
	adminModules.Put("/:id/update", controllers.UpdateModule)
	adminModules.Delete("/:id/delete", controllers.DeleteModule)

	// Admin Template Management
	adminTemplates := admin.Group("/templates")
	adminTemplates.Get("/", controllers.GetTemplateALL)
	adminTemplates.Post("/add", controllers.CreateTemplate)
	adminTemplates.Put("/:id/update", controllers.UpdateTemplate)

	// Admin User Module Management
	adminUserModules := admin.Group("/user_modules")
	adminUserModules.Get("/all", controllers.UserModuleFindAll)
	adminUserModules.Post("/add", controllers.UserModuleAddModule)
	adminUserModules.Delete("/delete", controllers.UserModuleDeleteModule)
}
