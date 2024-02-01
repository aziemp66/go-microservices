package log_routes

import (
	log_delivery "log-service/internal/delivery/log"

	"github.com/gofiber/fiber/v2"
)

func LogRoutes(router fiber.Router, logDelivery log_delivery.LogDelivery) {
	router.Get("", logDelivery.GetAllLog)
	router.Get("/:id", logDelivery.GetLogByID)
	router.Post("", logDelivery.CreateLog)
	router.Put("/:id", logDelivery.UpdateLogByID)
	router.Delete("/:id", logDelivery.DeleteLogByID)
	router.Post("/clear", logDelivery.ClearLog)
}
