package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Model
	CustomerID uint     `gorm:"customer_id"`
	Balance    float64  `gorm:"balance"`
	Customer   Customer `gorm:"foreignkey:customer_id;associationkey:id"`
}

func (a *Account) Validate() (err error) {
	if a.CustomerID == 0 {
		return fmt.Errorf("missing value for CustomerID")
	}

	if a.Balance < 0 {
		return fmt.Errorf("balance cannot be negative")
	}

	return nil
}

func (a *Account) BeforeSave() (err error) {
	if err = a.Validate(); err != nil {
		fmt.Println("can't save Account invalid data, err: ", err)

		return err
	}

	return nil
}
