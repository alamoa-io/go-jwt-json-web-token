package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-jwt-json-web-token/auth"
	"go-jwt-json-web-token/models"
)

func RefreshToken(c *gin.Context) {
	/*
		It is up for debate whether to send the refreshToken in the header or put it in the body. Either is fine since
		the data is sent together, but it is not defined in JWT. It is also possible to send the token as a Bear and the
		refreshToken as a Body. In terms of security, both the header and body are sent together, so they are the same.
		However, the header is processed first, so if you want to load it quickly, you can use the header.
	*/
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
