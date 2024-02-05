package http_server

import (
	http_error "mailer-service/internal/http/error"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewHttpServer() *fiber.App {
	app := fiber.New(fiber.Config{
		StrictRouting: true,
		CaseSensitive: true,
		ServerHeader:  "Logger",
		AppName:       "Logger Service v1.0",
		ErrorHandler:  errorHandler,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization,X-CSRF-TOKEN,Accept-Language",
	}))

	app.Get("", func(c *fiber.Ctx) error {
		c.SendString("App is running!!!")

		return nil
	})

	return app
}

func errorHandler(c *fiber.Ctx, err error) error {
	c.Response().Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	if httpError, ok := err.(*http_error.Error); ok {
		log.Error(httpError.Raw)
		return c.Status(httpError.Code).JSON(errResponse{
			Message: httpError.Error(),
		})
	} else {
		log.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(errResponse{
			Message: "Internal Server Error",
		})
	}
}
