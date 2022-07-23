package router

import (
	"github.com/gin-gonic/gin"
	"go_project/api/handler/v1/test"
	"go_project/pkg/middleware"
)


func APIRouter() *gin.Engine{
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.Use(middleware.LoggerToFile())
	liveV1 := router.Group("/api/test")
	{
		liveV1.GET("/find", test.Find)
	}
	return router
}
