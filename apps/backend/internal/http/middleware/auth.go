package middleware

import (
	"net/http"
	"strings"

	"github.com/SquadGO/squad-admin-panel/internal/http/helpers"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			helpers.JsonError(ctx, http.StatusUnauthorized, "Missing Authorization header")
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			helpers.JsonError(ctx, http.StatusUnauthorized, "Invalid Authorization header format")
			return
		}

		tokenString := parts[1]

		claims, err := helpers.ParseJWTToken(tokenString)

		if err != nil {
			helpers.JsonError(ctx, http.StatusUnauthorized, "Invalid token: "+err.Error())
			return
		}

		ctx.Set("steamID", claims.SteamID)
		ctx.Set("nickname", claims.NickName)
		ctx.Set("avatar", claims.Avatar)

		ctx.Next()
	}
}
