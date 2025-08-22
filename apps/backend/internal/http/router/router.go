package router

import (
	"fmt"
	"net/http"

	"github.com/SquadGO/squad-admin-panel/internal/http/handlers"
	"github.com/SquadGO/squad-admin-panel/internal/http/middleware"
	"github.com/SquadGO/squad-admin-panel/internal/service"
	"github.com/gin-gonic/gin"
)

func New(s *service.Service) *gin.Engine {
	r := gin.New()
	h := handlers.NewHandlers(s)

	r.Use(gin.Recovery())
	r.Use(middleware.Logger())

	{
		v1 := r.Group("api/v1")
		{
			authRouter := v1.Group("/auth")
			authRouter.GET("/steam", h.AuthHandler.Auth)
			authRouter.GET("/steam/success", h.AuthHandler.AuthSuccess)
		}

		{
			protected := v1.Group("/")
			protected.Use(middleware.JWTAuth())

			protected.GET("/user/:id")
		}
	}

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": map[string]any{
				"code":    http.StatusNotFound,
				"message": fmt.Sprintf("Resourse not found: %s", ctx.Request.URL.Path),
			},
		})
	})

	return r
}
