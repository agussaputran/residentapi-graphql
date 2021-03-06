package tables

import (
	"time"

	"gorm.io/gorm"
)

// Provinces table
type Provinces struct {
	ID           uint `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	ProvinceName string         `json:"province_name"`
	District     []Districts    `gorm:"ForeignKey:ProvinceID" json:"districts"`
}
