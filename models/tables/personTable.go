package tables

import (
	"time"

	"gorm.io/gorm"
)

// Persons Table
type Persons struct {
	ID                                                                                               uint `gorm:"primarykey" json:"id"`
	CreatedAt                                                                                        time.Time
	UpdatedAt                                                                                        time.Time
	DeletedAt                                                                                        gorm.DeletedAt `gorm:"index"`
	Nip, FullName, FirstName, LastName, BirthDate, BirthPlace, PhotoProfileURL, Gender, ZoneLocation string
	SubDistrictID                                                                                    uint
	OfficePersonLocation                                                                             []OfficePersonLocations `gorm:"ForeignKey:PersonID"`
}
