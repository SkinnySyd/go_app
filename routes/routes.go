package routes

import (
	"ginhello/controllers/logincontroller"
	"ginhello/controllers/todocontroller"
	"ginhello/controllers/usercontroller"
	"ginhello/middleware/authmiddleware"
	"ginhello/models"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	r := gin.Default()

	// Middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("currentUser", store))
	models.ConnectDatabase()

	// Routes
	todoRoutes := r.Group("/api/todo")
	{
		todoRoutes.Use(authmiddleware.RequiredLogin())
		todoRoutes.GET("/all", todocontroller.FindAllByUser)
		todoRoutes.GET("/:id", todocontroller.GetATodo)
		todoRoutes.POST("/add", todocontroller.AddTodo)
		todoRoutes.DELETE("/delete/:id", todocontroller.DeleteATodo)
		// Other todo routes...
	}

	r.GET("/api/user/users", usercontroller.Index)
	r.POST("/api/user/login", logincontroller.Login)
	r.POST("/api/user/register", logincontroller.Register)
	r.POST("/api/user/logout", logincontroller.Logout)

	// Other routes...

	r.Run()
}
