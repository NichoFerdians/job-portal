package models

type User struct {
	Name     string `json:"name"`
	UserName string `json:"user_name" gorm:"primaryKey"`
	Password string `json:"password"`
}

type LoginUserInput struct {
	UserName string `json:"user_name" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginUserResponse struct {
	AccessToken string `json:"access_token"`
	// User        UserResponse `json:"user"`
}

func (User) TableName() string {
	return "Users"
}
