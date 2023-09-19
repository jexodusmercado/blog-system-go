package storage

import (
	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) error {
	return db.AutoMigrate(
	)
}