package routes

import (
	"github.com/gofiber/fiber/v2"
	"pendaftaran-sidang/internal/controller"
	"pendaftaran-sidang/internal/middleware"
)

func SidangRoutes(router fiber.Router, controller controller.SidangController) {
	product := router.Group("/sidang")

	product.Post("/create", middleware.UserAuthentication(middleware.AuthConfig{
		Unauthorized: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	}), controller.Create)

	product.Patch("/update/:id", controller.Update)

}
