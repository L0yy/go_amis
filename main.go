package main

import (
	"amis_fiber/handles"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/django/v3"
)

func main() {
	engine := django.New("./static", ".tmpl")
	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins: "http://127.0.0.1:8080",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	app.Get("/", handles.HandleRootPath)
	app.Get("/index", handles.HandleRootPathIndex)

	app.Get("/GetData", handles.HandleGetConfig)

	log.Fatal(app.Listen(":8080"))
}
