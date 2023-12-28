package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-jwt-json-web-token/auth"
	"go-jwt-json-web-token/models"
)

func RefreshToken(c *gin.Context) {
	refreshTokenRequest := &struct {
		RefreshToken string `json:"refreshToken"`
	}{}
	if err := c.ShouldBindJSON(&refreshTokenRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}

	claims, err := auth.ValidateToken(refreshTokenRequest.RefreshToken, auth.RefreshTokenVerifyKey)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}

	if claims.IsBlackListToken() {
		c.JSON(400, gin.H{"Error": "is invalid token"})
		return
	}
	oldClaims := *claims
	defer oldClaims.SetBlackListToken()

	_, err = models.GetUserByEmail(claims.Email)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}

	refreshToken, err := claims.UpdateRefreshToken()
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}
	tokenClaim := &auth.Claim{Email: claims.Email}
	token, err := tokenClaim.GenerateToken()
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(200, &struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refreshToken"`
	}{
		Token:        token,
		RefreshToken: refreshToken,
	})
}
