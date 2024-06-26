package middleware

import (
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/authentication/dto"
	customErr "github.com/AkbarFikri/Mooistudio-Technical-Test/internal/pkg/error"
	jwt2 "github.com/AkbarFikri/Mooistudio-Technical-Test/internal/pkg/jwt"
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"strings"
)

func JwtUser() gin.HandlerFunc {

	return gin.HandlerFunc(func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			err := customErr.ErrorForbiddenAccess
			response.New(response.WithError(err)).SendAbort(c)
			return
		}

		if !strings.Contains(c.GetHeader("Authorization"), "Bearer") {
			err := customErr.ErrorUnauthorized
			response.New(response.WithError(err)).SendAbort(c)
			return
		}

		token, err := jwt2.VerifyTokenHeader(c)

		if err != nil {
			err := customErr.ErrorInvalidAccessToken
			response.New(response.WithError(err)).SendAbort(c)
			return
		} else {
			claims := token.Claims.(jwt.MapClaims)
			user := dto.UserTokenData{
				ID:    claims["id"].(string),
				Email: claims["email"].(string),
			}
			c.Set("user", user)
			c.Next()
		}
	})

}
