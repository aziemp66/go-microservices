package middleware

import (
	"errors"
	"mailer-service/internal/validation"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	id_translations "github.com/go-playground/validator/v10/translations/id"
	"github.com/gofiber/fiber/v2"
)

func ValidationMiddleware() (fiber.Handler, *validation.Validation) {
	v := validator.New()
	var trans ut.Translator
	return func(c *fiber.Ctx) error {
			en := en.New()
			id := id.New()
			t := ut.New(en, id)
			languages := c.GetReqHeaders()["Accept-Language"]
			var lang string
			if len(languages) == 0 {
				lang = "en"
			} else {
				lang = languages[0]
			}

			if lang == "en" {
				trans, isFound := t.FindTranslator("en")
				if !isFound {
					return errors.New("english language translator not found")
				}
				en_translations.RegisterDefaultTranslations(v, trans)
				v.RegisterTagNameFunc(func(field reflect.StructField) string {
					name := strings.SplitN(field.Tag.Get("en"), ",", 2)[0]
					// skip if tag key says it should be ignored
					if name == "-" {
						return ""
					}
					return name
				})
			} else if lang == "id" {
				trans, isFound := t.FindTranslator("id")
				if !isFound {
					return errors.New("bahasa indonesia translator not found")
				}
				id_translations.RegisterDefaultTranslations(v, trans)
				v.RegisterTagNameFunc(func(field reflect.StructField) string {
					name := strings.SplitN(field.Tag.Get("id"), ",", 2)[0]
					// skip if tag key says it should be ignored
					if name == "-" {
						return ""
					}
					return name
				})
			}

			return c.Next()
		}, &validation.Validation{
			Validator: v,
			Trans:     trans,
		}
}

func ApiMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Request().Header.Add(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

		return c.Next()
	}
}
