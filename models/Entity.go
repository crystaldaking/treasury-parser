package models

import "gorm.io/gorm"

type Entity struct {
	Uid       int `gorm:"primaryKey"`
	FirstName string
	LastName  string
	SndType   string
}

func FetchAll(db *gorm.DB) ([]Entity, error) {
	var entities []Entity

	if err := db.Find(&entities).Error; err != nil {
		return nil, err
	}

	return entities, nil
}
