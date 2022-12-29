package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

var versionProg = ""

func init() {

}

func main() {
	log.Println("Старт..")
	app := fiber.New()
	app.Static("/", "/home/arkadii/ProjectsGo2/fibergio/bin/main/html")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":8483")
}
