package tables

import "gorm.io/gorm"

// OfficePersonLocations table
type OfficePersonLocations struct {
	gorm.Model
	OfficeID uint
	PersonID uint
}
