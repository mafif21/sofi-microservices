package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
	"pendaftaran-sidang/internal/config"
	"pendaftaran-sidang/internal/controller"
	"pendaftaran-sidang/internal/repositories"
	"pendaftaran-sidang/internal/routes"
	"pendaftaran-sidang/internal/services"
)

func StartApp() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		},
	})
	app.Use(cors.New())

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"john":  "doe",
			"admin": "123456",
		},
	}))

	validator := validator.New()
	db := config.OpenConnection()

	repository := repositories.NewSidangRepository()
	service := services.NewSidangService(repository, db)
	sidangController := controller.NewSidangController(service, validator)

	api := app.Group("api")
	app.Static("/public/doc_ta", "./public/doc_ta")
	routes.SidangRoutes(api, sidangController)

	err := app.Listen(":" + os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
}
