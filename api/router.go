package api

import (
	"github.com/gin-gonic/gin"
	"github.com/solace06/AnraAssignment/internal"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	r := router.Group("/api/v1")

	internal.TaskRoutes(r)

	return router
}
