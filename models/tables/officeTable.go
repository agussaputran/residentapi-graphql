package tables

import "gorm.io/gorm"

// Offices table
type Offices struct {
	gorm.Model
	OfficeName           string                  `json:"office_name"`
	SubDistrictID        uint                    `json:"sub_district_id"`
	OfficePersonLocation []OfficePersonLocations `gorm:"ForeignKey:OfficeID"`
}
