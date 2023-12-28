package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"go-jwt-json-web-token/auth"
)

func Logout(c *gin.Context) {
	revokeTokenRequest := &struct {
		RefreshToken string `json:"refreshToken"`
	}{}
	if err := c.ShouldBindJSON(&revokeTokenRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}

	refreshTokenClaims, err := auth.ValidateToken(revokeTokenRequest.RefreshToken, auth.RefreshTokenVerifyKey)
	log.Print(err)
	if refreshTokenClaims != nil {
		refreshTokenClaims.SetBlackListToken()
	}

	// This code is duplicated with middlewares, so please move the common code elsewhere.
	authorizationHeader := c.Request.Header.Get("Authorization")
	bearerToken := strings.Split(authorizationHeader, "Bearer ")
	if len(bearerToken) != 2 {
		c.JSON(400, "invalid bearer token")
		c.Abort()
		return
	}
	token := strings.TrimSpace(bearerToken[1])
	tokenClaims, err := auth.ValidateToken(token, auth.TokenVerifyKey)
	log.Print(err)
	if tokenClaims != nil {
		tokenClaims.SetBlackListToken()
	}

	c.JSON(200, gin.H{"Message": "logout success"})
}
