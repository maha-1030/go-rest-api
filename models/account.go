package models

import "github.com/jinzhu/gorm"

type Accounts struct {
	gorm.Model
	CustomerID uint     `gorm:"customer_id"`
	Balance    float64  `gorm:"balance"`
	Customer   Customer `gorm:"foreignkey:customer_id;associationkey:id"`
}
