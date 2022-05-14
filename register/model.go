package register

import (
	"github.com/jinzhu/gorm"
	"go.quick.start/models"
)

func Models(db *gorm.DB) {
	models.GetUserModel(db)
}
