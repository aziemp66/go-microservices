package log_delivery

import (
	"errors"
	"fmt"
	"log-service/internal/data"
	http_server "log-service/internal/http"
	http_error "log-service/internal/http/error"
	"log-service/internal/models/request"

	"github.com/gofiber/fiber/v2"
)

func (d *logDelivery) CreateLog(c *fiber.Ctx) error {
	var req request.CreateLogEntry
	if err := c.BodyParser(&req); err != nil {
		return http_error.NewError(err, fiber.StatusBadRequest, "Gagal Parsing Body")
	}

	if err := d.validation.Validate(&req); err != nil {
		return http_error.NewError(err, fiber.StatusBadRequest, err.Error())
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

func (d *logDelivery) GetAllLog(c *fiber.Ctx) error {
	logs, err := d.models.LogEntry.All()
	if err != nil {
		return http_error.NewError(err, fiber.StatusBadRequest, "Gagal Mengambil Semua Log")
	}

	c.Status(fiber.StatusOK).JSON(http_server.Response{
		Message: "Berhasil Mengambil Semua Log",
		Value:   logs,
	})

	return nil
}

func (d *logDelivery) GetLogByID(c *fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return http_error.NewError(errors.New("params by the name of 'id' is empty"), fiber.StatusBadRequest, "ID tidak boleh kosong")
	}

	log, err := d.models.LogEntry.GetOne(id)
	if err != nil {
		return http_error.NewError(err, fiber.StatusNotFound, "Gagal mendapatkan Log")
	}

	c.Status(fiber.StatusOK).JSON(http_server.Response{
		Message: "Berhasil Mendapatkan Log",
		Value:   log,
	})

	return nil
}

// UpdateLogByID implements LogDelivery.
func (d *logDelivery) UpdateLogByID(c *fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return http_error.NewError(errors.New("params by the name of 'id' is empty"), fiber.StatusBadRequest, "ID tidak boleh kosong")
	}

	var req request.UpdateLogEntry
	if err := c.BodyParser(&req); err != nil {
		return http_error.NewError(err, fiber.StatusBadRequest, "Gagal Parsing Body")
	}

	if err := d.validation.Validate(&req); err != nil {
		return http_error.NewError(err, fiber.StatusBadRequest, err.Error())
	}

	d.models.LogEntry.ID = id
	d.models.LogEntry.Name = req.Name
	d.models.LogEntry.Data = req.Data

	result, err := d.models.LogEntry.Update()
	if err != nil {
		return http_error.NewError(err, fiber.StatusBadRequest, "Gagal Mengupdate Log")
	} else if result.ModifiedCount < 1 {
		return http_error.NewError(fmt.Errorf("log by the id of %s not found", id), fiber.StatusNotFound, "Tidak Menemukan Log Dengan ID Tersebut")
	}

	c.Status(fiber.StatusOK).JSON(http_server.Response{
		Message: "Berhasil Mengupdate Log",
	})

	return nil
}

// DeleteLogByID implements LogDelivery.
func (d *logDelivery) DeleteLogByID(c *fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return http_error.NewError(errors.New("params by the name of 'id' is empty"), fiber.StatusBadRequest, "ID tidak boleh kosong")
	}

	d.models.LogEntry.ID = id
	result, err := d.models.LogEntry.Delete()
	if err != nil {
		return http_error.NewError(err, fiber.StatusBadRequest, "Gagal Menghapus Log")
	} else if result.DeletedCount < 1 {
		return http_error.NewError(fmt.Errorf("log by the id of %s not found", id), fiber.StatusNotFound, "Tidak Menemukan Log Dengan ID Tersebut")
	}

	c.Status(fiber.StatusOK).JSON(http_server.Response{
		Message: "Berhasil Menghapus Log",
	})

	return nil
}

// ClearLog implements LogDelivery.
func (d *logDelivery) ClearLog(c *fiber.Ctx) error {
	err := d.models.LogEntry.DropCollection()
	if err != nil {
		return http_error.NewError(err, fiber.StatusBadRequest, "Gagal Menghapus Database")
	}

	c.Status(fiber.StatusOK).JSON(http_server.Response{
		Message: "Berhasil Membersihkan Log",
	})

	return nil
}
