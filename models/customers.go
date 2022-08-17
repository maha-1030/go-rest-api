package models

import (
	"github.com/jinzhu/gorm"
)

type Customer struct {
	gorm.Model
	Name        string `gorm:"name"`
	Age         int    `gorm:"age"`
	PhoneNumber string `gorm:"phone_number"`
	Gender      string `gorm:"gender"`
}
