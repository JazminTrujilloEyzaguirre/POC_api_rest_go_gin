package handler

import (
	"net/http"
	"root/httpd/resources/newsmembers"

	"github.com/gin-gonic/gin"
)

func NewMemberGet(members *newsmembers.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := members.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
