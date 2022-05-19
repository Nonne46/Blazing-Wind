package service

import (
	"els/internal/handlers"

	"github.com/gin-gonic/gin"
)

func NewHTTPHandler() *gin.Engine {
	r := gin.Default()

	search := r.Group("/search")
	search.GET("/posts", handlers.SearchPost)
	search.GET("/replies", handlers.SearchReplies)
	search.GET("/statuses", handlers.SearchStatuses)
	search.GET("/users", handlers.SearchUsers)

	return r
}
