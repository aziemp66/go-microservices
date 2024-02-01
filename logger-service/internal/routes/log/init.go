package log_routes

import (
	log_delivery "log-service/internal/delivery/log"

	"github.com/gofiber/fiber/v2"
)

func LogRoutes(router fiber.Router, logDelivery log_delivery.LogDelivery) {
	logRoutes := router.Group("/log")

	logRoutes.Post("", logDelivery.CreateLog)
}
