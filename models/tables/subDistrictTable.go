package tables

import "gorm.io/gorm"

// SubDistricts table
type SubDistricts struct {
	gorm.Model
	SubDistrictName string    `json:"sub_district_name"`
	DistrictID      uint      `json:"district_id"`
	Person          []Persons `gorm:"ForeignKey:SubDistrictID"`
}
