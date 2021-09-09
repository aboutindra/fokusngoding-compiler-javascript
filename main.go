package main

import (
	"fc-javascript/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var r router.Router

func init() {
	r = router.Router{}
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Post("/api/v1/compiler/javascript/run", r.Exec)
	app.Post("/api/v1/compiler/javascript/unit-test", r.Exec)

	app.Listen(":3500")
}
