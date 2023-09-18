package main

import (
	/* "github.com/gofiber/fiber/v2"
	"log" */
	"treasury-parser/services/parser"
	"treasury-parser/utils"
)

func main() {
	/* app := fiber.New()
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

	log.Fatal(app.Listen(":8080")) */

	db := utils.Init()
	data := parser.FetchData("https://www.treasury.gov/ofac/downloads/sdn.xml")
	parser.Import(db, parser.Parse(data))
}
