package handlers

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data": gin.H{
			"status":  "ok",
			"go":      runtime.Version(),
			"service": "baibaoxiang",
		},
	})
}
