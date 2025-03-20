package errors

import(
	"fmt"
	"strings"
	"github.com/go-playground/validator/v10"
	
)

type Error struct {
	Message string      `json:"message"`
	Err     string      `json:"error,omitempty"`
	Details interface{} `json:"details,omitempty"`
}


func NewError(message string, err error) *Error {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	return &Error{
		Message: message,
		Err:     errMsg,
	}
}

func NewValidationError(err error) *Error {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var details []map[string]string
		for _, validationError := range validationErrors {
			field := strings.ToLower(validationError.Field())
			details = append(details, map[string]string{
				"field":   field,
				"message": fmt.Sprintf("Field '%s' failed validation: %s", field, validationError.Tag()),
			})
		}
		return &Error{
			Message: "Validation failed",
			Err:     err.Error(),
			Details: details,
		}
	}
	return NewError("Validation failed", err)
}

