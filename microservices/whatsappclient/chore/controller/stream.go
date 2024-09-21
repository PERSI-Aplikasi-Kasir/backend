package controller

import (
	"backend/pkg/handler"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (c *whatsappController) RegisterStream(router *gin.Engine) {
	v1 := router.Group("/v1")

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

		if qrChan == nil {
			handler.Error(ctx, http.StatusBadRequest, "Device is already logged in")
			return
		}

		doneChan := ctx.Request.Context().Done()

		for {
			select {
			case qrImage, ok := <-qrChan:
				if !ok {
					return
				}
				if qrImage == nil {
					continue
				}

				_, err := ctx.Writer.Write(*qrImage)
				if err != nil {
					handler.Error(ctx, http.StatusInternalServerError, "Failed to write image data")
					return
				}

				flusher.Flush()

			case <-doneChan:
				log.Info().Msg("Client cancelled the QR login request")
				return
			}
		}
	})
}
