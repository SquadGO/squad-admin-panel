package middleware

import (
	"github.com/gin-gonic/gin"
)

func GothicProdiver() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		q := ctx.Request.URL.Query()
		q.Set("provider", "steam")
		ctx.Request.URL.RawQuery = q.Encode()
	}
}
