package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-jwt-json-web-token/auth"
)

func Logout(c *gin.Context) {
	revokeTokenRequest := &struct {
		// Since tokens expire quickly, it is sufficient to reject only refresh tokens that do not grant tokens.
		// In that case, it may be a good idea to make it unnecessary for the client to send a token.
		Token        string `json:"token"`
		RefreshToken string `json:"refreshToken"`
	}{}
	if err := c.ShouldBindJSON(&revokeTokenRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}

	tokenClaims, err := auth.ValidateToken(revokeTokenRequest.Token, auth.TokenVerifyKey)
	log.Print(err)
	if tokenClaims != nil {
		tokenClaims.SetBlackListToken()
	}
	refreshTokenClaims, err := auth.ValidateToken(revokeTokenRequest.RefreshToken, auth.RefreshTokenVerifyKey)
	log.Print(err)
	if refreshTokenClaims != nil {
		refreshTokenClaims.SetBlackListToken()
	}

	c.JSON(200, gin.H{"Message": "logout success"})
}
