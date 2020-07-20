package models

import "github.com/jinzhu/gorm"

// User ...
type User struct {
	gorm.DB
	Name     string `gorm:"not null"`
	Password string `json:"-" gorm:"not null"`
	Salt     string `json:"-" gorm:"default:'******';not null"`
	Mobile   string `gorm:"type:varchar(50);unique_index;not null"`
	Code     string `gorm:"type:varchar(100);unique_index;not null"`
	Enabled  int32  `gorm:"type:tinyint(1);default:1;not null"`
}
