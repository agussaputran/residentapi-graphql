package tables

import (
	"time"

	"gorm.io/gorm"
)

// QueueEmail Table
type QueueEmail struct {
	ID        uint `gorm:"primarykey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Handled   bool
	To        string
	Cc        string
	Subject   string
	Message   string
}
