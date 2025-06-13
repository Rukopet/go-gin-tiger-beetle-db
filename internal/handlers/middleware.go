package handlers

import (
	"net/http"

	"finance-api/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggingMiddleware() gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			logrus.WithFields(logrus.Fields{
				"status":     param.StatusCode,
				"method":     param.Method,
				"path":       param.Path,
				"ip":         param.ClientIP,
				"latency":    param.Latency,
				"user_agent": param.Request.UserAgent(),
			}).Info("HTTP Request")
			return ""
		},
	})
}

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		logrus.WithField("error", recovered).Error("Panic recovered")

		errorResponse := models.ErrorResponse{
			Error:   "Internal Server Error",
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}

		c.JSON(http.StatusInternalServerError, errorResponse)
		c.Abort()
	})
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func handleServiceError(c *gin.Context, err error) {
	logrus.WithError(err).Error("Service error")

	errorResponse := models.ErrorResponse{
		Error:   "Service Error",
		Message: err.Error(),
		Code:    http.StatusInternalServerError,
	}

	c.JSON(http.StatusInternalServerError, errorResponse)
}
