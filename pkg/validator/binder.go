package validator

import (
	"backend/pkg/handler"
	"io"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func BindBody(c *gin.Context, req interface{}) (ok bool) {
	if err := c.ShouldBindJSON(req); err != nil {
		if err == io.EOF {
			handler.Error(c, http.StatusBadRequest, "Invalid JSON data: unexpected end of JSON input")
			log.Error().Msg("Invalid JSON data: unexpected end of JSON input")
			return false
		}
		handler.Error(c, http.StatusBadRequest, err.Error())
		log.Error().Err(err).Msg("Invalid JSON data")
		return false
	}

	if err := requestValidator(c, req, "body"); err != nil {
		return false
	}

	return true
}

func BindBodies(c *gin.Context, obj interface{}) (ok bool) {
	v := reflect.ValueOf(obj).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i).Addr().Interface()
		if err := c.ShouldBindJSON(field); err != nil {
			if err == io.EOF {
				handler.Error(c, http.StatusBadRequest, "Invalid JSON data: unexpected end of JSON input")
				log.Error().Msg("Invalid JSON data: unexpected end of JSON input")
				return false
			}
			handler.Error(c, http.StatusBadRequest, err.Error())
			log.Error().Err(err).Msg("Invalid JSON data")
			return false
		}

		if err := requestValidator(c, field, "body"); err != nil {
			return false
		}
	}

	return true
}

func BindParam(c *gin.Context, req interface{}) (ok bool) {
	if err := c.ShouldBindQuery(req); err != nil {
		handler.Error(c, http.StatusBadRequest, err.Error())
		log.Error().Err(err).Msg("Invalid query data")
		return false
	}

	if err := requestValidator(c, req, "param"); err != nil {
		return false
	}

	return true
}

func BindParams(c *gin.Context, obj interface{}) (ok bool) {
	v := reflect.ValueOf(obj).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i).Addr().Interface()
		if err := c.ShouldBindQuery(field); err != nil {
			handler.Error(c, http.StatusBadRequest, err.Error())
			log.Error().Err(err).Msg("Invalid query data")
			return false
		}

		if err := requestValidator(c, field, "param"); err != nil {
			return false
		}
	}

	return true
}

func BindUri(c *gin.Context, req interface{}) (ok bool) {
	if err := c.ShouldBindUri(req); err != nil {
		handler.Error(c, http.StatusBadRequest, err.Error())
		log.Error().Err(err).Msg("Invalid uri data")
		return false
	}

	if err := requestValidator(c, req, "uri"); err != nil {
		return false
	}

	return true
}

func BindUris(c *gin.Context, obj interface{}) (ok bool) {
	v := reflect.ValueOf(obj).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i).Addr().Interface()
		if err := c.ShouldBindUri(field); err != nil {
			handler.Error(c, http.StatusBadRequest, err.Error())
			log.Error().Err(err).Msg("Invalid uri data")
			return false
		}

		if err := requestValidator(c, field, "uri"); err != nil {
			return false
		}
	}

	return true
}
