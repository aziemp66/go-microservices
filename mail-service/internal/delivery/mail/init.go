package mail_delivery

import "mailer-service/internal/validation"

type delivery struct {
	Validation validation.Validation
}

func NewMailDelivery(validation *validation.Validation) MailDelivery {
	return &delivery{
		Validation: *validation,
	}
}
