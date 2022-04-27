package main

import "github.com/gofiber/fiber/v2"

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awasome api!")
}

func main(){
	app := fiber.New()

	app.Get("/api", welcome)

	log.Fatal(app.Listen(":3000"))
}