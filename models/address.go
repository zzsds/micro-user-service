package models

import (
	"github.com/jinzhu/gorm"
)

// Address ...
type Address struct {
	gorm.DB
	UserID    int32   `gorm:"index;not null"`
	Name      string  `gorm:"type:varchar(50);not null"`
	Country   string  `gorm:"type:varchar(50);default:'中国'"`
	Province  string  `gorm:"type:varchar(50);not null"`
	City      string  `gorm:"type:varchar(50);not null"`
	Area      string  `gorm:"type:varchar(50);not null"`
	Address   string  `gorm:"type:varchar(500);not null"`
	Mobile    string  `gorm:"type:varchar(20);not null"`
	Tel       string  `gorm:"type:varchar(20)"`
	Email     string  `gorm:"type:varchar(50)"`
	IDCard    string  `gorm:"type:varchar(30)"`
	PostCode  string  `gorm:"type:varchar(10)"`
	Longitude float64 `gorm:"type:decimal(11,7);default:0"`
	Latitude  float64 `gorm:"type:decimal(11,7);default:0"`
	IDJust    string  `gorm:"type:varchar(250)"`
	IDBack    string  `gorm:"type:varchar(250)"`
	Default   int32   `gorm:"type:tinyint(1);default:0;not null"`
	User      *User
}
