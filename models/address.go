package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

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

func (a *Address) Validate() (err error) {
	if a.CustomerID == 0 {
		return fmt.Errorf("missing value for customerID")
	}

	if a.Line1 == "" {
		return fmt.Errorf("missing value for line1")
	}

	if a.City == "" {
		return fmt.Errorf("missing value for city")
	}

	if a.State == "" {
		return fmt.Errorf("missing value for state")
	}

	if a.PinCode == "" {
		return fmt.Errorf("missing value for pincode")
	}

	return nil
}

func (a *Address) BeforeSave() (err error) {
	if err = a.Validate(); err != nil {
		fmt.Println("can't save Address invalid data, err: ", err)

		return err
	}

	return nil
}
