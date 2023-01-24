package helpers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator"
)

// ValidateStruct validates a struct using validator package.
func ValidateStruct(data interface{}) error {
	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		switch err := err.(type) {
		case validator.ValidationErrors:
			var fields []string
			for _, field := range err {
				fields = append(fields, strings.ToLower(field.Field()))
			}
			return fmt.Errorf("field validation error, missing fields: %v", fields)
		default:
			return fmt.Errorf("field validation error: %w", err)
		}
	}
	return nil
}
