package controller

import "github.com/gofiber/fiber/v2"

type SidangController interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	FindByUser(ctx *fiber.Ctx) error
}
