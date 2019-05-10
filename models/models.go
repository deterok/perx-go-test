package models

import "github.com/jinzhu/gorm"

var prepareModels []interface{}

func InitAllModels(db *gorm.DB) error {
	for _, model := range prepareModels {

		if err := db.AutoMigrate(model).Error; err != nil {
			return err
		}
	}

	return nil
}
