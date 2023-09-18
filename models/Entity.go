package models

type Entity struct {
	Uid       int `gorm:"primaryKey"`
	FirstName string
	LastName  string
	SndType   string
}
