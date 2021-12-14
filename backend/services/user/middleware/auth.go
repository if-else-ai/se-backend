package middleware

import (
	"log"
	"net/http"

	"kibby/user/auth"

	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := auth.TokenValid(c.Request); err != nil {
			log.Println(err.Error())
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		userID, err := auth.ExtractIDFromToken(c.Request)
		if err != nil {
			log.Println(err.Error())
			c.Status(http.StatusUnprocessableEntity)
			c.Abort()
			return
		}

		c.Set("userID", userID)

		c.Next()
	}
}
