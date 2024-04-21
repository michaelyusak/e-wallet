package middleware

import (
	"os"
	"time"

	"e-wallet/constants"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger(log *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery

		ctx.Next()

		if raw != "" {
			path = path + "?" + raw
		}

		statusCode := ctx.Writer.Status()

		requestId, exist := ctx.Get(string(constants.RequestId))
		if !exist {
			requestId = ""
		}

		entry := log.WithFields(logrus.Fields{
			"request_id":  requestId,
			"latency":     time.Since(start),
			"method":      ctx.Request.Method,
			"status_code": statusCode,
			"path":        path,
		})

		file, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer file.Close()

		log.SetOutput(file)

		logEntry := entry.Info
		if statusCode >= 400 && statusCode < 600 {
			log.Error(entry)
			return
		}

		logEntry("request processed")
	}
}
