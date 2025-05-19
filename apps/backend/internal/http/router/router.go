package router

import (
	"fmt"
	"net/http"

	"github.com/SquadGO/squad-admin-panel/internal/http/middleware"
	"github.com/SquadGO/squad-admin-panel/internal/service"
	"github.com/gin-gonic/gin"
)

func New(s service.Service) *gin.Engine {
	r := gin.New()
	//handlers := handlers.NewHandlers(s)

	r.Use(gin.Recovery())
	r.Use(middleware.Logger())

	// r.GET("/auth", handlers.UserHandlers.Auth)
	// r.POST("/reg", handlers.UserHandlers.Reg)

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": map[string]any{
				"code":    http.StatusNotFound,
				"message": fmt.Sprintf("Resourse not found: %s", ctx.Request.URL.Path),
			},
		})
	})

	// r.GET("/user", h.UserHandlers.Get)
	// r.POST("/auth", h.UserHandlers.Authorization)
	// r.POST("/reg", h.UserHandlers.Registration)

	// authorized := r.Group("/")
	// authorized.Use(middleware.AuthRequired())
	// {
	// 	authorized.DELETE("/todo/:id", h.TodoHandlers.Delete)
	// 	authorized.GET("/todo/:id", h.TodoHandlers.GetTodo)
	// 	authorized.GET("/todos", h.TodoHandlers.GetTodos)
	// 	authorized.POST("/todo", h.TodoHandlers.Create)
	// 	authorized.PATCH("/todo", h.TodoHandlers.Update)
	// }

	return r
}
