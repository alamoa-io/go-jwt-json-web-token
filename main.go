package main

import (
	"github.com/gin-gonic/gin"

	"go-jwt-json-web-token/controllers"
	"go-jwt-json-web-token/middlewares"
)

func main() {
	r := gin.Default()
	api := r.Group("/api")
	{
		v1 := api.Group("/v1/")
		{
			v1.POST("/login", controllers.Login)
			v1.POST("/signup", controllers.Signup)
			v1.POST("/refresh_token", controllers.RefreshToken)
			v1.GET("/profile", middlewares.JwtTokenVerifier(), controllers.Profile)
			v1.POST("/logout", middlewares.JwtTokenVerifier(), controllers.Logout)
		}
	}
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
