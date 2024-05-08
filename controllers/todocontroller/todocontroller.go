package todocontroller

import (
	"ginhello/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func FindAll(c *gin.Context) {

	var todo []models.Todo

	models.DB.Find(&todo)
	c.JSON(http.StatusOK, gin.H{"todo": todo})

}
func FindAllByUser(c *gin.Context) {

	var todo []models.Todo

	session := sessions.Default(c)
	userID := session.Get("userID")
	if userID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	userIDint, ok := userID.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
		return
	}
	todo, err := models.GetTodosByUserID(userIDint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get todo list", "userID": userIDint})
		return
	}

	//models.DB.Find(&todo)
	c.JSON(http.StatusOK, gin.H{"todo": todo})

}

func AddTodo(c *gin.Context) {
	session := sessions.Default(c)
	var todo models.Todo

	userIDint, ok := session.Get("userID").(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
		return
	}

	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo.UserID = userIDint
	if err := models.DB.Create(&todo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)

}

func GetATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo models.Todo
	session := sessions.Default(c)
	userID := session.Get("userID")
	userIDstr := userID.(int64)
	if userID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	if err := models.DB.Where("id = ? AND user_id = ?", id, userIDstr).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	// if err := models.DB.Where("id = ?", id).First(&todo).Error; err != nil {
	// 	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
	// 	return
	// }

	c.JSON(http.StatusOK, todo)
}

func DeleteATodo(c *gin.Context) {
	id := c.Params.ByName("id")

	var todo models.Todo

	//models.DB.Delet(&todo,id)
	if err := models.DB.Where("id = ?", id).Delete(&todo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"todo with id:" + id: "deleted"})
}

func UpdateATodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")

	// Get the todo from the database
	if err := models.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	// Bind JSON data to the todo struct
	if err := c.BindJSON(&todo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Update the todo in the database
	if err := models.DB.Save(&todo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}
