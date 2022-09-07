package main

import (
	"root/httpd/handler"
	"root/httpd/resources/newsmembers"

	"github.com/gin-gonic/gin"
)

func main() {
	memberslocaldb := newsmembers.New()
	r := gin.Default()
	r.GET("/ping", handler.GetPing())
	r.GET("/team/:name", handler.GetTeam) // parameters in path
	r.GET("/welcome", handler.GetWelcome) // parameters in querystring \\ /welcome?firstname=Jaz&lastname=Trujillo
	r.POST("/", handler.PostSimple)       // body and err

	r.GET("/member", handler.NewMemberGet(memberslocaldb))
	r.POST("/member", handler.NewMemberPost(memberslocaldb))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
