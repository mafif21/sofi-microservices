package routes

import (
	"github.com/gofiber/fiber/v2"
	"pendaftaran-sidang/internal/controller"
)

func SidangRoutes(router fiber.Router, controller controller.SidangController) {
	product := router.Group("/sidang")

	product.Post("/create", controller.Create)
}
