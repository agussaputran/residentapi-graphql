package tables

import "gorm.io/gorm"

// OfficePersonLocations table
type OfficePersonLocations struct {
	gorm.Model
	OfficeID uint `json:"office_id"`
	PersonID uint `json:"person_id"`
}
