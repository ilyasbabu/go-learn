package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ilyasbabu/e-com/database"
	"github.com/ilyasbabu/e-com/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.JSON("Welcome")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)
	// users
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)
	// products
	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Put("/api/products/:id", routes.UpdateProduct)
	app.Delete("/api/products/:id", routes.DeleteProduct)
	// orders
	app.Post("/api/orders", routes.CreateOrder)
	app.Get("/api/orders", routes.GetOrders)

}

func main() {
	database.ConnectDb()
	app := fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(":8000"))
}
