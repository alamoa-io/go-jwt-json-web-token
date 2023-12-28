package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"

	"go-jwt-json-web-token/auth"
)

func JwtTokenVerifier() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.Request.Header.Get("Authorization")
		bearerToken := strings.Split(authorizationHeader, "Bearer ")
		if len(bearerToken) != 2 {
			c.JSON(400, "invalid bearer token")
			c.Abort()
			return
		}
		token := strings.TrimSpace(bearerToken[1])
		claims, err := auth.ValidateToken(token, auth.TokenVerifyKey)
		if err != nil {
			c.JSON(400, err.Error())
			c.Abort()
			return
		}

		// If you want to include the access token in the blacklist when you log out, you can set it as follows,
		// but this is not often done. Also, JWT can be allowed without using a database etc. If you want to use it,
		// it is better to use Redis because it is fast. It depends on the application, but in many cases it is possible
		// to retain the token by simply adjusting the Token's Expire rather than rejecting it with blacklist.
		if claims.IsBlackListToken() {
			c.JSON(400, "invalid access")
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Next()
	}
}
