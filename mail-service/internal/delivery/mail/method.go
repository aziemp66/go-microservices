package mail_delivery

import (
	http_server "mailer-service/internal/http"
	http_error "mailer-service/internal/http/error"
	"mailer-service/internal/model/request"
	mail_usecase "mailer-service/internal/usecase/mail"

	"github.com/gofiber/fiber/v2"
)

func (d *delivery) SendEmail(c *fiber.Ctx) error {
	var req request.SendEmail
	if err := c.BodyParser(&req); err != nil {
		return http_error.NewError(err, fiber.StatusBadRequest, "Gagal Parsing Body")
	}

	if err := d.Validation.Validate(&req); err != nil {
		return err
	}

	err := d.MailUsecase.SendSMTPMessage(mail_usecase.Message{
		From:    *req.From,
		To:      *req.To,
		Subject: *req.Subject,
		Data:    *req.Message,
	})
	if err != nil {
		return http_error.NewError(err, fiber.StatusBadRequest, "Email Gagal Dikirim")
	}

	return c.Status(fiber.StatusOK).JSON(http_server.Response{
		Message: "Berhasil Mengirim Email",
	})
}
