package config

import (
	"os"

	"github.com/Yan-pi/go-opportunities-api/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite Migration Error", err)
		return nil, err
	}

	return db, err
}
