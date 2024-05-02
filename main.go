package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		isLoggedIn := session.Get("isLoggedIn")

		if isLoggedIn != true {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func main() {

	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.LoadHTMLGlob("templates/*")

	r.GET("/", authMiddleware(), func(c *gin.Context) {
		c.String(http.StatusOK, "hello world v3")
	})

	r.GET("/login", func(c *gin.Context) {
		session := sessions.Default(c)
		isLoggedIn := session.Get("isLoggedIn")
		if isLoggedIn == true {
			c.Redirect(http.StatusFound, "/")
			return
		}
		c.HTML(http.StatusOK, "login.tmpl", gin.H{})
	})

	r.POST("/login", func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")

		// TODO: validate user's credentials
		if email == "user@mail" && password == "password" {
			// Successful login
			session := sessions.Default(c)
			session.Set("isLoggedIn", true)
			session.Save()
			c.Redirect(http.StatusFound, "/")
		} else {
			// Invalid credentials
			c.HTML(http.StatusBadRequest, "login.tmpl", gin.H{
				"errorMessage": "Invalid email or password",
			})
		}
	})

	r.Run()
}
