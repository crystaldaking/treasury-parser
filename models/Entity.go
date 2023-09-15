package models

import (
	"gorm.io/gorm"
)

type Entity struct {
	gorm.Model
	Id        int `gorm:"primaryKey"`
	FirstName string
	LastName  string
	SndType   bool
}
