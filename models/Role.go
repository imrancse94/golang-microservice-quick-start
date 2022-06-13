package models

import (
	"go.quick.start/Helper"
)

type Role struct {
	ID     int    `gorm:"primary_key" json:"id"`
	Title  string `gorm:"type:varchar(100);not null" json:"title"`
	Status int    `gorm:"type:tinyint(4);not null" json:"status"`
}

func GetRoles() Helper.ModelResponse {
	var roles []Role
	err := DB.Find(&roles).Error
	var resp Helper.ModelResponse
	resp.Data = roles
	resp.Message = "User found successfully"
	if err != nil {
		resp.Data = nil
		resp.Message = "User not found"
	}
	return resp
}
