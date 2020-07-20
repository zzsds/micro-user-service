package models

import "github.com/jinzhu/gorm"

// Profile ...
type Profile struct {
	gorm.DB
	UserID   int32  `gorm:"unique_index;not null"`
	Province string `gorm:"type:varchar(50);not null"`
	City     string `gorm:"type:varchar(50);not null"`
	Area     string `gorm:"type:varchar(50);not null"`
}
