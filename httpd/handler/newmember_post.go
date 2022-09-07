package handler

import (
	"net/http"
	"root/httpd/resources/newsmembers"

	"github.com/gin-gonic/gin"
)

type newsmembersPostRequest struct {
	Name       string `json:"name"`
	OriginTeam string `json:"originTeam"`
}

func NewMemberPost(member *newsmembers.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := newsmembersPostRequest{}
		c.Bind(&requestBody)

		item := newsmembers.Item{
			Name:       requestBody.Name,
			OriginTeam: requestBody.OriginTeam,
		}
		member.Add(item)
		c.JSON(http.StatusOK, map[string]string{
			"status":  "Succeded",
			"message": "Member was added Succesfull",
		})
		c.Status(http.StatusNoContent)
	}
}
