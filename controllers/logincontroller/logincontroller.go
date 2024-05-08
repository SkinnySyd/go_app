package logincontroller

import (
	"ginhello/models"
	"net/http"

	"github.com/gin-contrib/sessions"

	// "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Find user by username
	foundUser, err := models.GetUserByUsername(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
		return
	}

	// Check user exists and password is correct
	if foundUser != nil && foundUser.CheckPassword(user.Password) {
		session := sessions.Default(c)
		session.Set("isLoggedIn", true)
		session.Set("userID", foundUser.Id)
		session.Save()
		c.JSON(http.StatusOK, gin.H{"message": "Login successful", "userID": foundUser.Id, "user ID": session.Get("userID")})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
	}

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
