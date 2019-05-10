package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Code struct {
	gorm.Model
	Code   string `gorm:"unique"`
	UsedAt *time.Time
}

func init() {
	prepareModels = append(prepareModels, &Code{})
}
