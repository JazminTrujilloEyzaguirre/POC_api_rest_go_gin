package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWelcome(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
	c.String(http.StatusOK, "Hola %s %s", firstname, lastname)
}
