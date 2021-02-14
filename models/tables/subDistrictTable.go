package tables

import (
	"time"

	"gorm.io/gorm"
)

// SubDistricts table
type SubDistricts struct {
	ID              uint `gorm:"primarykey" json:"id"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	SubDistrictName string         `json:"sub_district_name"`
	DistrictID      uint           `json:"district_id"`
	Person          []Persons      `gorm:"ForeignKey:SubDistrictID"`
}
