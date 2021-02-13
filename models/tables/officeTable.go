package tables

import "gorm.io/gorm"

// Offices table
type Offices struct {
	gorm.Model
	Name                 string
	SubDistrictID        uint
	OfficePersonLocation []OfficePersonLocations `gorm:"ForeignKey:OfficeID"`
}
