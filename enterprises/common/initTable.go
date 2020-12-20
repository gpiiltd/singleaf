package common

import (
	"singleaf/enterprises/models"

	"github.com/jinzhu/gorm"
)

// InitTable use for drop/delete table if is exists
func InitTable(db *gorm.DB) {
	db.DropTableIfExists(&models.Enterprises{})
}
