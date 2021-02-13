package tables

import "gorm.io/gorm"

// SubDistricts table
type SubDistricts struct {
	gorm.Model
	Name       string
	DistrictID uint
	Person     []Persons `gorm:"ForeignKey:SubDistrictID"`
}
