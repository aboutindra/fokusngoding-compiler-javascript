package main

import (
	"fc-javascript/router"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

var r router.Router

func init() {
	r = router.Router{}
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Post("/api/v1/compiler/javascript/run", r.Exec)

	app.Listen(4000)
}
