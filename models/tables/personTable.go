package tables

import (
	"time"

	"gorm.io/gorm"
)

// Persons Table
type Persons struct {
	ID                   uint `gorm:"primarykey" json:"id"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt          `gorm:"index"`
	Nip                  string                  `json:"nip"`
	FullName             string                  `json:"full_name"`
	FirstName            string                  `json:"first_name"`
	LastName             string                  `json:"last_name"`
	BirthDate            string                  `json:"birth_date"`
	BirthPlace           string                  `json:"birth_place"`
	PhotoProfileURL      string                  `json:"photo_profile_url"`
	Gender               string                  `json:"gender"`
	ZoneLocation         string                  `json:"zone_location"`
	SubDistrictID        uint                    `json:"sub_district_id"`
	OfficePersonLocation []OfficePersonLocations `gorm:"ForeignKey:PersonID"`
}
