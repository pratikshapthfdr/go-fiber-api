package api

import (
	"example.com/go-fiber-api/database"
	"github.com/gofiber/fiber/v2"
)
  
 var (
	inputs = map[string]database.Input{}
 )
 func createInput(c *fiber.Ctx) error {
	input := new(database.Input)
  
	err := c.BodyParser(&input)
  
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err.Error(), 	
		})
  
	}
  
	//input.Name = New().string()
	//fmt.Scanln(&input.Name)
	inputs[input.Name] = *input
    
	c.Status(200).JSON(&fiber.Map{ 
        "responce" :  "Received Request" ,
		"input_data": input,
	}) 
  
	return nil
 }

func inputData(c *fiber.Ctx) error {
	c.Status(200).JSON(&fiber.Map{
		"inputs": inputs,
	})
	return nil
 }

func SetupRoute() *fiber.App {
	app := *fiber.New() 
	app.Get("/api/v1/inputs/:id", inputData)
	//app.Post("/api/v1/inputs/responce", "Received Request")
	app.Post("/api/v1/inputs", createInput)
   // app.Get("/api/v1/articles/:id", readArticle)
    
return &app
}