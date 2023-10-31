package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/srikanthankamtw/todo-rest/model"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Task{})
	db.AutoMigrate(&model.User{})
}
