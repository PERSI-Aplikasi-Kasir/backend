package controller

import (
	"backend/microservices/whatsappclient/chore/entity"
	"backend/microservices/whatsappclient/chore/interfaces"
	"backend/microservices/whatsappclient/chore/service"
	"backend/pkg/handler"
	"backend/pkg/validator"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mau.fi/whatsmeow"
)

type whatsappController struct {
	service interfaces.WhatsappService
}

func NewWhatsappController(client *whatsmeow.Client) *whatsappController {
	var (
		controller = new(whatsappController)
		service    = service.NewWhatsappService(client)
	)
	controller.service = service

	return controller
}

func (c *whatsappController) Register(router *gin.Engine) {
	v1 := router.Group("/v1")

	v1.GET("/whatsapp/is-login", func(ctx *gin.Context) {
		ok, err := c.service.CheckDevice()
		if err != nil {
			handler.Error(ctx, http.StatusNotFound, err.Error())
			return
		}

		handler.Success(ctx, http.StatusOK, "Success checking whatsapp login", gin.H{
			"is_login": ok,
		})
	})

	v1.GET("/whatsapp/login-qr", func(ctx *gin.Context) {
		flusher, ok := ctx.Writer.(http.Flusher)
		if !ok {
			handler.Error(ctx, http.StatusInternalServerError, "Streaming unsupported")
			return
		}

		ctx.Writer.Header().Set("Content-Type", "image/png")
		ctx.Writer.Header().Set("Cache-Control", "no-cache")
		ctx.Writer.Header().Set("Connection", "keep-alive")

		qrChan, err := c.service.GetLoginQR()
		if err != nil {
			handler.Error(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		for qrImage := range qrChan {
			if qrImage == nil {
				continue
			}

			_, err := ctx.Writer.Write(*qrImage)
			if err != nil {
				handler.Error(ctx, http.StatusInternalServerError, "Failed to write image data")
				return
			}

			flusher.Flush()
		}

		ctx.Writer.WriteHeader(http.StatusOK)
		flusher.Flush()
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
}
