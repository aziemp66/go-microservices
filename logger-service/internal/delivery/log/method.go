package log_delivery

import (
	"log-service/internal/data"
	http_server "log-service/internal/http"
	http_error "log-service/internal/http/error"
	"log-service/internal/models/request"

	"github.com/gofiber/fiber/v2"
)

func (d *LogDelivery) CreateLog(c *fiber.Ctx) error {
	var req request.CreateLogEntry
	if err := c.BodyParser(&req); err != nil {
		return http_error.NewError(err, fiber.StatusBadRequest, "Gagal Parsing Body")
	}

	if err := d.validation.Validate(&req); err != nil {
		return err
	}

	err := d.models.LogEntry.Insert(data.LogEntry{
		Name: req.Name,
		Data: req.Data,
	})
	if err != nil {
		return http_error.NewError(err, fiber.StatusBadRequest, "Gagal Menambahkan Log")
	}

	c.Status(fiber.StatusCreated).JSON(http_server.Response{
		Message: "Berhasil Menambahkan Log",
		Value:   req,
	})

	return nil
}
