package tables

import (
	"time"

	"gorm.io/gorm"
)

// Offices table
type Offices struct {
	ID                   uint `gorm:"primarykey" json:"id"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt          `gorm:"index"`
	OfficeName           string                  `json:"office_name"`
	SubDistrictID        uint                    `json:"sub_district_id"`
	OfficePersonLocation []OfficePersonLocations `gorm:"ForeignKey:OfficeID"`
}
