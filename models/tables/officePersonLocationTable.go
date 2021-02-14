package tables

import (
	"time"

	"gorm.io/gorm"
)

// OfficePersonLocations table
type OfficePersonLocations struct {
	ID        uint `gorm:"primarykey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	OfficeID  uint           `json:"office_id"`
	PersonID  uint           `json:"person_id"`
}
