package models

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
)

// User ...
type User struct {
	gorm.Model
	Name     string     `validate:"required" gorm:"not null"`
	Password string     `json:"-"`
	Salt     string     `json:"-" gorm:"default:'******';not null"`
	Mobile   string     `validate:"required" gorm:"type:varchar(50);unique_index;not null"`
	Email    string     `gorm:"type:varchar(100);not null"`
	Nickname string     `gorm:"type:varchar(50)"`
	Realname string     `gorm:"type:varchar(50)"`
	Code     string     `gorm:"type:varchar(100);unique_index;not null"`
	Source   string     `validate:"required" gorm:"index;type:varchar(100);not null"`
	Birthday *time.Time `gorm:"type:datetime"`
	Enabled  int32      `gorm:"type:tinyint(1);default:1"`
}

// BeforeCreate ...
func (u *User) BeforeCreate(scope *gorm.Scope) error {
	var randCode string
	if u.Code != "" {
		randCode = u.Code
	} else {
		randCode = GenerateCode(6)
	}
	salt := EncodeMD5(randCode)
	scope.SetColumn("salt", salt)
	if u.Password != "" {
		pass, _ := EncodeSalt(u.Password, salt)
		scope.SetColumn("password", pass)
	}
	return nil
}

// EncodeMD5 md5 encryption
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

// EncodeSalt ...
func EncodeSalt(password, salt string) (string, error) {
	dk, err := scrypt.Key([]byte(password), []byte(salt), 1<<15, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
		return password, err
	}
	return base64.StdEncoding.EncodeToString(dk), nil
}

// GenerateCode ...
func GenerateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

// ValidateMobile ...
func ValidateMobile(mobile string) bool {
	ok, _ := regexp.MatchString(`^((\+[0-9]\d{10,12})|1[1-9]\d{9})$`, mobile)
	return ok
}
