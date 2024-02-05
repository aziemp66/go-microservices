package mail_delivery

import (
	http_server "mailer-service/internal/http"

	"github.com/gofiber/fiber/v2"
)

func (d *delivery) SendEmail(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(http_server.Response{
		Message: "Berhasil Mengirim Email",
	})
}
