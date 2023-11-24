package main

import (
	"sipd-pembiayaan/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	app.Get("/api/pts", handler.Getpts)
	app.Get("/api/tpb", handler.Gettpb)
	app.Get("/api/sts", handler.Getsts)

	app.Listen(":8081")

}
