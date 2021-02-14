package tables

import "gorm.io/gorm"

// Districts table
type Districts struct {
	gorm.Model
	DistrictName string
	ProvinceID   uint
	SubDistrict  []SubDistricts `gorm:"ForeignKey:DistrictID"`
}
