package tables

import "gorm.io/gorm"

// QueueEmail Table
type QueueEmail struct {
	gorm.Model
	Handled bool
	To      string
	Cc      string
	Subject string
	Message string
}
