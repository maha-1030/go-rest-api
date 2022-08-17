package models

import "github.com/jinzhu/gorm"

type Address struct {
	gorm.Model
	CustomerID uint     `gorm:"customer_id"`
	Line1      string   `gorm:"line1"`
	Line2      string   `gorm:"line2"`
	City       string   `gorm:"city"`
	State      string   `gorm:"state"`
	PinCode    string   `gorm:"pin_code"`
	Customer   Customer `gorm:"foreignkey:customer_id;associationkey:id"`
}
