package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// FormatValidationError formats validation errors into a user-friendly format
func FormatValidationError(err error) gin.H {
    if validationErrors, ok := err.(validator.ValidationErrors); ok {
        formattedErrors := make(map[string]string)
        for _, fieldError := range validationErrors {
            formattedErrors[strings.ToLower(fieldError.Field())] = fieldError.Tag() 
        }
        return gin.H{"errors": formattedErrors}
    }
    return gin.H{"error": "Invalid input"}
}
