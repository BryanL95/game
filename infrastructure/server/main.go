package main

import (
	"log"

	"github.com/BryanL95/game/infrastructure/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":9001"))
}
