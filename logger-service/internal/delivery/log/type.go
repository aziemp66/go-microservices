package log_delivery

import "github.com/gofiber/fiber/v2"

type LogDelivery interface {
	CreateLog(c *fiber.Ctx) error
	GetAllLog(c *fiber.Ctx) error
	GetLogByID(c *fiber.Ctx) error
	UpdateLogByID(c *fiber.Ctx) error
	DeleteLogByID(c *fiber.Ctx) error
	ClearLog(c *fiber.Ctx) error
}
