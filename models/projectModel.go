package models

import (
	"time"
	 "gorm.io/gorm"
)

type Project struct {
	gorm.Model

	ID	uint
	Name string
	Hidden bool
	Icon string
	Deadline *time.Time
}