package storage

import (
	"fmt"
	"log"

	"github.com/jexodusmercado/blog-system-go/internal/conf"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Dial(conf conf.Config) (*gorm.DB) {
	DSN := conf.DbHost

	fmt.Println(DSN)
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{
		// DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalln("database connection failed")
		return nil
	}

	err = InitMigration(db)
	if err != nil {
		log.Fatalln("database migration failed")
		return nil
	}

	return db
}