package controllers

import (
	"github.com/gin-gonic/gin"

	"go-jwt-json-web-token/models"
)

func Signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}
	err = user.CreateUser()
	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"Message": "signup success"})
}
