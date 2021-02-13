package tables

import "gorm.io/gorm"

// Districts table
type Districts struct {
	gorm.Model
	Name        string
	ProvinceID  uint
	SubDistrict []SubDistricts `gorm:"ForeignKey:DistrictID"`
}
