package controllers

import (
	"github.com/gin-gonic/gin"

	"go-jwt-json-web-token/models"
)

func Profile(c *gin.Context) {
	email, _ := c.Get("email")
	user, err := models.GetUserByEmail(email.(string))
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(200, user)
}
