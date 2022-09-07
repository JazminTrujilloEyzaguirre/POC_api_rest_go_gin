package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTeam(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hola %s  Team", name)
}
