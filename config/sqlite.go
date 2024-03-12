package config

import (
	"os"

	"github.com/Yan-pi/go-opportunities-api/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	return db, err
}
