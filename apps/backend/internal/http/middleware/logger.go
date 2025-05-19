package middleware

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log := slog.With(
			slog.String("method", ctx.Request.Method),
			slog.String("path", ctx.Request.URL.Path),
			slog.String("remote_addr", ctx.Request.RemoteAddr),
			slog.String("user_agent", ctx.Request.UserAgent()),
		)

		log.Info("REQUEST:")
	}
}
