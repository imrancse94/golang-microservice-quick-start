package models

import (
	"github.com/jinzhu/gorm"
	"go.quick.start/Helper"
)

var DB *gorm.DB

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `gorm:"type:varchar(100);not null" json:"name"`
	Email    string `gorm:"type:varchar(100);not null;unique" json:"email"`
	Password string `gorm:"type:varchar(100);not null" json:"-"` // hidden field
}

// RegisterUserInput Embed the user struct properties
type RegisterUserInput struct {
	User
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

func GetUserByEmail(email string) Helper.ModelResponse {
	var user User
	err := DB.Where("email = ?", email).First(&user).Error
	var resp Helper.ModelResponse
	resp.Data = user
	resp.Message = "User found successfully"
	if err != nil {
		resp.Data = nil
		resp.Message = "User not found"
	}
	return resp
}

func GetUserById(id int) Helper.ModelResponse {
	var user User
	err := DB.First(&user, id).Error
	var resp Helper.ModelResponse
	resp.Data = user
	resp.Message = "User found successfully"
	if err != nil {
		resp.Data = nil
		resp.Message = "User not found"
	}
	return resp
}

func GetUserModel(db *gorm.DB) {
	DB = db
	var user User
	DB.AutoMigrate(&user)
}
