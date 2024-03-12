package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	logger *Logger
)

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
