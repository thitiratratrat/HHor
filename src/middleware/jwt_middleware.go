package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/thitiratratrat/hhor/src/service"
)

func AuthorizeJWT(role service.Role) gin.HandlerFunc {
	return func(context *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := context.GetHeader("Authorization")
		id := context.Param("userid")

		if !(len(BEARER_SCHEMA) >= 0 && len(BEARER_SCHEMA) <= len(authHeader)) {
			context.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := service.JWTServiceHandler().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)

			if string(role) != claims["role"] {
				context.AbortWithStatus(http.StatusUnauthorized)
			}

			if len(id) > 0 && claims["id"] != id {
				context.AbortWithStatus(http.StatusUnauthorized)
			}

		} else {
			logrus.Error(err)
			context.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
