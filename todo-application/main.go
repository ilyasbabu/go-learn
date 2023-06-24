package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

type UpdateResponse struct {
	Updated Todo   `json:"updated"`
	Todos   []Todo `json:"todos"`
	Respose string `json:"response"`
}

func main() {
	fmt.Println("Hello World")

	app := fiber.New()
	todos := []Todo{}

	app.Get("/check", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Post("/api/todo", func(c *fiber.Ctx) error {
		todo := &Todo{}

		err := c.BodyParser(todo)
		fmt.Printf("%T", err)
		if err != nil {
			return err
		}
		// title := c.FormValue("title")
		// body := c.FormValue("body")
		todo.ID = len(todos) + 1
		// todo.Title = title
		// todo.Body = body
		todos = append(todos, *todo)
		return c.JSON(todos)
	})

	app.Get("/api/todo", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	app.Patch("/api/todo/:id/done", func(c *fiber.Ctx) error {
		// id := c.AllParams()
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.SendStatus(404)
			// return c.Status(404).SendString("Id not found")
		}
		fmt.Println(id)

		for index, todo := range todos {
			if todo.ID == id {
				todos[index].Done = true
				res := &UpdateResponse{}
				res.Updated = todo
				res.Todos = todos
				res.Respose = "Update"
				return c.JSON(res)
			}
		}

		return c.JSON("Invalid id")
	})

	log.Fatal(app.Listen(":8000"))
}
