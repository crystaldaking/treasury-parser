package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"treasury-parser/utils"
)

func main() {
	app := fiber.New()
	utils.Init()

	app.Post("/update", func(ctx *fiber.Ctx) error {
		return ctx.SendString("implement update route")
	})

	app.Get("/state", func(ctx *fiber.Ctx) error {
		return ctx.SendString("implement state route")
	})

	app.Get("/get_names", func(ctx *fiber.Ctx) error {
		return ctx.SendString("implement get names route")
	})

	log.Fatal(app.Listen(":8080"))
}
