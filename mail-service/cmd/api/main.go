package main

import (
	mail_delivery "mailer-service/internal/delivery/mail"
	http_server "mailer-service/internal/http"
	"mailer-service/internal/http/middleware"
	mail_router "mailer-service/internal/router/mail"
)

func main() {
	srv := http_server.NewHttpServer()

	validationMidleware, validation := middleware.ValidationMiddleware()

	api := srv.Group("/api")
	api.Use(validationMidleware)

	mailDelivery := mail_delivery.NewMailDelivery(validation)
	mail_router.NewMailRouter(api.Group("/mail"), mailDelivery)

	err := srv.Listen(":80")
	if err != nil {
		panic(err)
	}
}
