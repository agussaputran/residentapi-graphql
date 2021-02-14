package tables

import "gorm.io/gorm"

// Districts table
type Districts struct {
	gorm.Model
	DistrictName string         `json:"district_name"`
	ProvinceID   uint           `json:"province_id"`
	SubDistrict  []SubDistricts `gorm:"ForeignKey:DistrictID"`
}
