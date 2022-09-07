package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPing() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"message": "pong",
		})
		// c.JSON(http.StatusOK, gin.H{
		// 	"message": "pong",
		// })
	}
}
