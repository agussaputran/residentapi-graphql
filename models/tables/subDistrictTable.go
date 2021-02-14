package tables

import "gorm.io/gorm"

// SubDistricts table
type SubDistricts struct {
	gorm.Model
	SubDistrictName string
	DistrictID      uint
	Person          []Persons `gorm:"ForeignKey:SubDistrictID"`
}
