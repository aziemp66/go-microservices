package mail_router

import (
	mail_delivery "mailer-service/internal/delivery/mail"

	"github.com/gofiber/fiber/v2"
)

func NewMailRouter(router fiber.Router, mailDelivery mail_delivery.MailDelivery) {
	router.Post("", mailDelivery.SendEmail)
}
