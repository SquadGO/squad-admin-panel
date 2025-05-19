package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func jsonError(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, gin.H{
		"error": map[string]any{
			"code":    statusCode,
			"message": message,
		},
	})
}

func isValidJSON(ctx *gin.Context, err error) {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, err := range validationErrors {
			if err.Field() == "Email" && err.Tag() == "email" {
				jsonError(ctx, http.StatusBadGateway, "Invalid email format")
				return
			}

			if err.Field() == "Password" && err.Tag() == "min" {
				jsonError(ctx, http.StatusBadGateway, "Min password length 8")
				return
			}

			if err.Field() == "Password" && err.Tag() == "max" {
				jsonError(ctx, http.StatusBadGateway, "Max password length 32")
				return
			}
		}

		jsonError(ctx, http.StatusBadGateway, "Validation failed")
		return
	} else {
		jsonError(ctx, http.StatusBadGateway, "Missing required params")
	}
}
