package main

import (
	"ginhello/controllers/logincontroller"
	"ginhello/controllers/todocontroller"
	"ginhello/controllers/usercontroller"
	"ginhello/models"

	//"ginhello/routes"

	//"ginhello/controllers/usercontroller"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
func main() {

	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	models.ConnectDatabase()
	todoroutes := r.Group("/api/todo")
	{
		todoroutes.Use(requiredLogin())
		todoroutes.GET("/all", todocontroller.FindAllByUser)
		todoroutes.GET("/:id", todocontroller.GetATodo)
		todoroutes.POST("/add", todocontroller.AddTodo)
		todoroutes.DELETE("/delete/:id", todocontroller.DeleteATodo)
	}

	r.GET("/api/user/users", usercontroller.Index)
	r.POST("/api/user/login", logincontroller.Login)
	r.POST("/api/user/register", logincontroller.Register)
	r.POST("/api/user/logout", logincontroller.Logout)

	r.GET("/user/login-required", func(c *gin.Context) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You need to be logged in"})
	})
	//
	//
	// r.LoadHTMLGlob("templates/*")

	// r.GET("/", authMiddleware(), func(c *gin.Context) {
	// 	c.String(http.StatusOK, "hello world v4")
	// })

	// r.GET("/login", func(c *gin.Context) {
	// 	session := sessions.Default(c)
	// 	isLoggedIn := session.Get("isLoggedIn")
	// 	if isLoggedIn == true {
	// 		c.Redirect(http.StatusFound, "/")
	// 		return
	// 	}
	// 	c.HTML(http.StatusOK, "login.tmpl", gin.H{})
	// })

	// r.POST("/login", func(c *gin.Context) {
	// 	email := c.PostForm("email")
	// 	password := c.PostForm("password")

	// 	// TODO: validate user's credentials
	// 	if email == "user@mail" && password == "password" {
	// 		// Successful login
	// 		session := sessions.Default(c)
	// 		session.Set("isLoggedIn", true)
	// 		session.Save()
	// 		c.Redirect(http.StatusFound, "/")
	// 	} else {
	// 		// Invalid credentials
	// 		c.HTML(http.StatusBadRequest, "login.tmpl", gin.H{
	// 			"errorMessage": "Invalid email or password",
	// 		})
	// 	}
	// })

	r.Run()
}
