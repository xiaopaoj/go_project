package test

import (
	"github.com/gin-gonic/gin"
	"go_project/internal/service"
	"net/http"
)

func Find(c *gin.Context) {
	t, err := service.Svc.Test().Find(c)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "find!",
		"data":    t,
	})
}
