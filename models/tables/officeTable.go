package tables

import "gorm.io/gorm"

// Offices table
type Offices struct {
	gorm.Model
	OfficeName           string
	SubDistrictID        uint
	OfficePersonLocation []OfficePersonLocations `gorm:"ForeignKey:OfficeID"`
}
