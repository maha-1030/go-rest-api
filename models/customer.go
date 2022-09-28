package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Customer struct {
	gorm.Model
	Name        string    `gorm:"name"`
	Age         int       `gorm:"age"`
	PhoneNumber string    `gorm:"phone_number"`
	Gender      string    `gorm:"gender"`
	Accounts    []Account `gorm:"foreignKey:customer_id;references:id"`
	Addresses   []Address `gorm:"foreignKey:customer_id;references:id"`
}

func (c *Customer) Validate() (err error) {
	if c.Name == "" {
		return fmt.Errorf("missing value for name")
	}

	if c.Age == 0 {
		return fmt.Errorf("missing value for age")
	}

	if c.PhoneNumber == "" {
		return fmt.Errorf("missing value for phone number")
	}

	if c.Gender == "" {
		return fmt.Errorf("missing value for gender")
	}

	if c.Gender != "M" && c.Gender != "F" {
		return fmt.Errorf("invalid value for gender")
	}

	return nil
}

func (c *Customer) BeforeSave() (err error) {
	if err = c.Validate(); err != nil {
		fmt.Println("can't save Customer invalid data, err: ", err)

		return err
	}

	return nil
}
