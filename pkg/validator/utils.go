package validator

import (
	"backend/pkg/handler"
	"errors"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var tagMessages = map[string]string{
	"required":       "Kolom ini wajib diisi",
	"required-param": "Parameter ini wajib diisi",
	"email":          "Email tidak valid",
	"max":            "Melebihi panjang maksimum",
	"min":            "Tidak memenuhi panjang minimum",
	"len":            "Panjang tidak valid",
	"lte":            "Melebihi nilai maksimum",
	"gte":            "Tidak memenuhi nilai minimum",
}

func parseTagMessage(tag string) string {
	if message, ok := tagMessages[tag]; ok {
		return message
	}
	return tag
}

func requestValidator(c *gin.Context, req interface{}, validatorType string) error {
	var modelTag string
	var errorMessage string

	if validatorType == "body" {
		modelTag = "json"
		errorMessage = "Permintaan body tidak valid"
	} else if validatorType == "param" {
		modelTag = "form"
		errorMessage = "Parameter query tidak valid"
	} else if validatorType == "uri" {
		modelTag = "uri"
		errorMessage = "Parameter URI tidak valid"
	} else {
		panic("Jenis validator tidak valid")
	}

	val := validator.New()
	val.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get(modelTag), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	if err := val.Struct(req); err != nil {
		var valErrors validator.ValidationErrors
		if errors.As(err, &valErrors) {
			errors := make([]handler.ApiError, len(valErrors))
			for i, valError := range valErrors {
				errorTag := valError.Tag()
				if errorTag == "required" && validatorType == "param" {
					errorTag = "required-param"
				}
				errors[i] = handler.ApiError{
					Field:   valError.Field(),
					Message: parseTagMessage(errorTag),
				}
			}
			handler.Error(c, http.StatusBadRequest, errorMessage, errors...)
		}
		return err
	}
	return nil
}
