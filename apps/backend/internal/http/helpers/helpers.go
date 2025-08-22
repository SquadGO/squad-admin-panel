package helpers

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	NickName string `json:"nickname"`
	Avatar   string `json:"avatar"`
	SteamID  string `json:"steamID"`
	jwt.StandardClaims
}

func GenerateJWTToken(nickname, avatar, steamID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		NickName: nickname,
		Avatar:   avatar,
		SteamID:  steamID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return tokenString, err
}

func ParseJWTToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*Claims); ok {
		return claims, nil
	} else {
		return nil, errors.New("Parse token error")
	}
}

func JsonError(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, gin.H{
		"error": map[string]any{
			"code":    statusCode,
			"message": message,
		},
	})
}

func IsValidJSON(ctx *gin.Context, err error) {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, err := range validationErrors {
			if err.Field() == "Email" && err.Tag() == "email" {
				JsonError(ctx, http.StatusBadGateway, "Invalid email format")
				return
			}

			if err.Field() == "Password" && err.Tag() == "min" {
				JsonError(ctx, http.StatusBadGateway, "Min password length 8")
				return
			}

			if err.Field() == "Password" && err.Tag() == "max" {
				JsonError(ctx, http.StatusBadGateway, "Max password length 32")
				return
			}
		}

		JsonError(ctx, http.StatusBadGateway, "Validation failed")
		return
	} else {
		JsonError(ctx, http.StatusBadGateway, "Missing required params")
	}
}
