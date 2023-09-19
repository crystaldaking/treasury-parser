package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"treasury-parser/controllers"
	"treasury-parser/utils"
)

func main() {
	app := fiber.New()
	db := utils.Init()
	handler := controllers.NewBaseHandler(db)

	app.Post("/update", handler.Update)

	app.Get("/state", handler.State)

	app.Get("/get_names", func(ctx *fiber.Ctx) error {
		return ctx.SendString("implement get names route")
	})

	log.Fatal(app.Listen(":8080"))
}
