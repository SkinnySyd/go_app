package logincontroller

import (
	"ginhello/models"
	"ginhello/services/loginservice"
	"net/http"

	"github.com/gin-contrib/sessions"

	// "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Login handles user login
func Login(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Authenticate user
	user, err := loginservice.Authenticate(loginData.Username, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Set session if login is successful
	saveSession(c, user)
	// Return user information
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    user,
	})
}

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Hash the password
	if err := user.SetPassword(user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create user
	if err := models.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	isLoggedIn := session.Get("isLoggedIn")
	if isLoggedIn == true {
		session.Clear()
		session.Save()
		c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
	}

}

func saveSession(c *gin.Context, user *models.User) {
	session := sessions.Default(c)
	session.Set("userID", user.Id)
	session.Set("isLoggedIn", true)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
}
