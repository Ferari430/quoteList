package req

import (
	"log"
	"quoteList/internal/payload"

	"github.com/go-playground/validator/v10"
)

func IsValid(payload payload.Quote) error {
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		log.Println("Validator bad")

		return err

	}
	return nil
}
