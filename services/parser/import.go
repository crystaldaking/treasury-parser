package parser

import (
	"treasury-parser/models"

	"gorm.io/gorm"
)

func Import(db *gorm.DB, list sndList) {
	for _, value := range list.SndEntry {
		entry := models.Entity{
			Uid:       value.Uid,
			FirstName: value.FirstName,
			LastName:  value.LastName,
			SndType:   value.SdnType,
		}

		db.FirstOrCreate(&entry)
	}
}
