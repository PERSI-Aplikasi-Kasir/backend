package controller

import (
	"backend/microservices/whatsappclient/chore/entity"
	"backend/microservices/whatsappclient/chore/interfaces"
	"backend/microservices/whatsappclient/chore/service"
	"backend/microservices/whatsappclient/config"
	"backend/pkg/handler"
	"backend/pkg/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type whatsappController struct {
	service interfaces.WhatsappService
}

func NewWhatsappController() *whatsappController {
	client := config.GetClient()

	return &whatsappController{
		service: service.NewWhatsappService(client),
	}
}

func (c *whatsappController) Register(router *gin.Engine) {
	v1 := router.Group("/v1")

	v1.GET("/whatsapp/is-login", func(ctx *gin.Context) {
		isLogin, err := c.service.CheckDevice()
		if err != nil {
			handler.Error(ctx, http.StatusNotFound, err.Error())
			return
		}

		handler.Success(ctx, http.StatusOK, "Success checking whatsapp login", gin.H{
			"is_login": isLogin,
		})
	})

	v1.POST("/whatsapp", func(ctx *gin.Context) {
		var req entity.MessageSend
		if !validator.BindBody(ctx, &req) {
			return
		}

		err := c.service.SendMessage(&req)
		if err != nil {
			handler.Error(ctx, http.StatusNotFound, err.Error())
			return
		}

		handler.Success(ctx, http.StatusOK, "Success sending whatsapp message", nil)
	})

	v1.GET("/whatsapp/reset", func(ctx *gin.Context) {
		err := c.service.ResetLoggedDevice()
		if err != nil {
			handler.Error(ctx, http.StatusNotFound, err.Error())
			return
		}

		handler.Success(ctx, http.StatusOK, "Success resetting whatsapp device", nil)
	})
}
