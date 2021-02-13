package tables

import "gorm.io/gorm"

// Provinces table
type Provinces struct {
	gorm.Model
	Name     string
	District []Districts `gorm:"ForeignKey:ProvinceID"`
}
