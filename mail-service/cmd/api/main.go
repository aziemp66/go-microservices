package main

import (
	mail_delivery "mailer-service/internal/delivery/mail"
	http_server "mailer-service/internal/http"
	"mailer-service/internal/http/middleware"
	mail_router "mailer-service/internal/router/mail"
	mail_usecase "mailer-service/internal/usecase/mail"
	"os"
	"strconv"
)

func main() {
	srv := http_server.NewHttpServer()

	validationMidleware, validation := middleware.ValidationMiddleware()

	api := srv.Group("/api")
	api.Use(validationMidleware)

	port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		panic("Email Port is not integer")
	}
	mailUsecase := mail_usecase.NewEmailUsecase(
		os.Getenv("MAIL_DOMAIN"),
		os.Getenv("MAIL_HOST"),
		port,
		os.Getenv("MAIL_USERNAME"),
		os.Getenv("MAIL_PASSWORD"),
		os.Getenv("MAIL_ENCRYPTION"),
		os.Getenv("MAIL_FROM_ADDRESS"),
		os.Getenv("MAIL_FROM_NAME"),
	)
	mailDelivery := mail_delivery.NewMailDelivery(validation, mailUsecase)
	mail_router.NewMailRouter(api.Group("/mail"), mailDelivery)

	err = srv.Listen(":80")
	if err != nil {
		panic(err)
	}
}
