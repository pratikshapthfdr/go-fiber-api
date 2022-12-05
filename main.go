package main

import (
	"example.com/go-fiber-api/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := api.SetupRoute()
 
	app.Post("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Go Fiber API")
	})
  
	  app.Listen(":3000")
}
