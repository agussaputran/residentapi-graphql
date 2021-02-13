package tables

import "gorm.io/gorm"

// Persons Table
type Persons struct {
	gorm.Model
	Nip, FullName, FirstName, LastName, BirthDate, BirthPlace, PhotoProfileURL, Gender, ZoneLocation string
	SubDistrictID                                                                                    uint
	OfficePersonLocation                                                                             []OfficePersonLocations `gorm:"ForeignKey:PersonID"`
}
