package storage

import (
	"github.com/jinzhu/gorm"
	"github.com/kkirsche/cronmon/model"
)

func DatabaseSetup(db *gorm.DB) {
	db.AutoMigrate(model.Host{}, model.Task{})
}
