package jwt

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"Blog_Backend/pkg/e"
	"Blog_Backend/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		Headers := c.Request.Header
		token := ""
		if _, ok := Headers["Token"]; !ok {
			code = e.INVALID_PARAMS
		} else {
			token = Headers["Token"][0]
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}
		c.Next()
	}
}
