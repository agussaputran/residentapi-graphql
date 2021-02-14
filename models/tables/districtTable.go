package tables

import (
	"time"

	"gorm.io/gorm"
)

// Districts table
type Districts struct {
	ID           uint `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	DistrictName string         `json:"district_name"`
	ProvinceID   uint           `json:"province_id"`
	SubDistrict  []SubDistricts `gorm:"ForeignKey:DistrictID"`
}
