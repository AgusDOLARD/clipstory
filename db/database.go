package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func Init() error {
	var err error
	Db, err = gorm.Open(sqlite.Open("/tmp/clipstory.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return fmt.Errorf("Database Init Error: %w", err)
	}
	return nil
}
