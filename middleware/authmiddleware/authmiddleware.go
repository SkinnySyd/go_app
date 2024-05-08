package authmiddleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	//"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func requiredLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		isLoggedIn := session.Get("isLoggedIn")
		if isLoggedIn != true {
			//c.Redirect(http.StatusFound, "/api/user/login-required")
			c.JSON(http.StatusForbidden, gin.H{"StatusForbidden": "you need to login"})

			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}
