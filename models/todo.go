package models

//import "time"

type Todo struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"type:varchar(300)" json:"title"`
	Description string `gorm:"type:varchar(300)" json:"description"`
	UserID      int64  `json:"user_id"`
	// CreatedAt   *time.Time
	// UpdatedAt   *time.Time
}

func GetTodosByUserID(userID int64) ([]Todo, error) {
	var todos []Todo
	if err := DB.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}
