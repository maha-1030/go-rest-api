package customer

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/maha-1030/go-rest-api/models"
	"github.com/maha-1030/go-rest-api/store"
)

type customer struct {
	db *gorm.DB
}

func NewCustomer(db *gorm.DB) store.Customer {
	return &customer{
		db: db,
	}

}

func (c *customer) Get() (customers []models.Customer, err error) {
	if c.db == nil {
		fmt.Println("Err: Database connection is not established")

		return nil, fmt.Errorf("database connection is not established")
	}

	res := c.db.Find(&customers)
	if res.Error != nil {
		fmt.Println("Error while retrieving customers data, err: ", res.Error)

		return nil, res.Error
	}

	return customers, nil
}
