package mail_delivery

import (
	mail_usecase "mailer-service/internal/usecase/mail"
	"mailer-service/internal/validation"
)

type delivery struct {
	Validation  validation.Validation
	MailUsecase mail_usecase.MailUsecase
}

func NewMailDelivery(validation *validation.Validation, mailUsecase mail_usecase.MailUsecase) MailDelivery {
	return &delivery{
		Validation:  *validation,
		MailUsecase: mailUsecase,
	}
}
