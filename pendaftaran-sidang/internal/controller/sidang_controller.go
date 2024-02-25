package controller

import "github.com/gofiber/fiber/v2"

type SidangController interface {
	Create(ctx *fiber.Ctx) error
}
