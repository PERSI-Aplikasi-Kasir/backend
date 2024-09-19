package logger

import (
	"backend/pkg/env"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w ResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func DiscordLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		if env.DiscordWebhookUrl == "" {
			return
		}
		startTime := time.Now()

		w := &ResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = w
		c.Next()
		latency := time.Since(startTime)

		method := c.Request.Method
		path := c.Request.URL.Path
		urlParams := c.Request.URL.Query()
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		timestamp := startTime.Format("15:04:05")

		var urlParamsEncoded string
		if len(urlParams) != 0 {
			urlParamsEncoded = "?" + urlParams.Encode()
		}
		pathAndQueryParams := path + urlParamsEncoded

		var logMessage string

		if statusCode >= 200 && statusCode < 300 {
			logMessage = fmt.Sprintf(
				"> `%s` **[%s]** %s\n> **%d** | %s | %v",
				timestamp, method, pathAndQueryParams, statusCode, clientIP, latency,
			)
		} else {
			responseBody := w.body.String()

			var prettyResponseBody bytes.Buffer
			json.Indent(&prettyResponseBody, []byte(responseBody), "", "  ")

			logMessage = fmt.Sprintf(
				"> `%s` **[%s]** %s\n> **%d** | %s | %v\n**Response:**\n```json\n%s\n```",
				timestamp, method, pathAndQueryParams, statusCode, clientIP, latency, prettyResponseBody.String(),
			)
		}

		go sendToDiscord(logMessage)
	}
}

func sendToDiscord(message string) {
	webhookBody := map[string]string{
		"content": message,
	}
	jsonBody, err := json.Marshal(webhookBody)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", env.DiscordWebhookUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending log to Discord:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		fmt.Println("Error: non-200 status code received from Discord:", resp.Status)
	}
}
