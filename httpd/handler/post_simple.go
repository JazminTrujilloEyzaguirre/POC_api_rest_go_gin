package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PostSimple(c *gin.Context) {
	statusCode := c.Request.Method
	if statusCode == "POST" {
		body := c.Request.Body
		value, err := ioutil.ReadAll(body)
		if err != nil {
			fmt.Println(err.Error())
		}
		tn := time.Now()

		c.JSON(http.StatusOK, gin.H{
			"status":  "succeded",
			"message": string(value),
			"time":    tn,
		})
	}
}
