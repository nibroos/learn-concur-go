// main.go
package main

import (
    "github.com/gofiber/fiber/v2"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func main() {
    app := fiber.New()

    users := []User{
        {ID: 1, Name: "John"},
        {ID: 2, Name: "Janea"},
    }

    app.Get("/users", func(c *fiber.Ctx) error {
        return c.JSON(users)
    })

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hellw, sadesz!")
    })

    app.Listen(":4000")
}
