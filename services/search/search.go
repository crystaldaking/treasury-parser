package search

import (
	"gorm.io/gorm"
	"strings"
	"treasury-parser/models"
)

const strongCondition = "strong"
const weakCondition = "weak"

func DoSearch(db *gorm.DB, condition string, searchText string) []models.Entity {
	var entities []models.Entity

	switch strings.ToLower(condition) {
	case strongCondition:
		db.Where("LOWER(first_name || ' ' || last_name) ILIKE ?", "%"+strings.ToLower(searchText)+"%").Find(&entities)
		return entities
	case weakCondition:
		searchArray := strings.Split(searchText, " ")
		query := db
		for _, searchTerm := range searchArray {
			q := db.Where("LOWER(first_name || ' ' || last_name) ILIKE ?", "%"+strings.ToLower(searchTerm)+"%")
			query = query.Or(q)
		}
		query.Find(&entities)
		return entities
	}

	return nil
}
