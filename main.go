package main

import (
	"cybercampus_module/configs"
	//"cybercampus_module/mock"
	"cybercampus_module/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {

    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })
    
    configs.MongoConnect()

	//mock.InitMock()
	
    routes.InitRoutes(app)

    printRoutes(app)

    app.Listen(configs.LoadEnv("PORT"))


}

func printRoutes(app *fiber.App) {
	routes := app.Stack() 
	fmt.Println("Registered Routes:")
	for method, routeStack := range routes {
		for _, route := range routeStack {
			fmt.Printf("%d %s\n", method, route.Path)
		}
	}
}