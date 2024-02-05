package mail_delivery

import "github.com/gofiber/fiber/v2"

type MailDelivery interface {
	SendEmail(c *fiber.Ctx) error
}
