package routes

import (
	"github.com/BryanL95/game/interface/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(r *fiber.App) {
	r.Use(cors.New())

	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the best game after TEKKEN")
	})

	r.Post("/new-fight", controllers.CreateFight)
}
