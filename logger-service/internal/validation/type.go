package validation

import (
	"errors"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Validate struct {
	Validator *validator.Validate
	Trans     ut.Translator
}

func (v *Validate) Validate(data any) error {
	// type User struct {
	// 	Username string `validate:"required"`
	// 	Tagline  string `validate:"required,lt=10"`
	// 	Tagline2 string `validate:"required,gt=1"`
	// }

	// user := User{
	// 	Username: "Joeybloggs",
	// 	Tagline:  "This tagline is way too long.",
	// 	Tagline2: "1",
	// }

	err := v.Validator.Struct(data)
	if err != nil {

		// translate all error at once
		errs := err.(validator.ValidationErrors)

		// returns a map with key = namespace & value = translated error
		// NOTICE: 2 errors are returned and you'll see something surprising
		// translations are i18n aware!!!!
		// eg. '10 characters' vs '1 character'
		// type val struct {
		// 	Key   string
		// 	Value string
		// }
		// fmt.Println(errs.Translate(v.Trans))
		// type valError struct {
		// 	Key   string
		// 	Error interface{}
		// }
		// var structErr valError
		for _, e := range errs {
			// can translate each error one at a time.
			return errors.New(e.Translate(v.Trans))
		}
	}
	return nil
}
