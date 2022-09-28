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

func (c *customer) GetAll() (customers []models.Customer, err error) {
	res := c.db.Model(&models.Customer{}).Preload("Accounts").Preload("Addresses").Find(&customers)
	if res.Error != nil {
		fmt.Println("Error while retrieving customers data, err: ", res.Error)

		return nil, res.Error
	}

	return customers, nil
}

func (c *customer) Create(customer *models.Customer) (newCustomer *models.Customer, err error) {
	if res := c.db.Create(customer); res.Error != nil {
		fmt.Println("Error while creating the new customer, err: ", err)

		return nil, err
	}

	return customer, nil
}

func (c *customer) Get(id int) (customer *models.Customer, err error) {
	customer = &models.Customer{}

	if res := c.db.Model(&models.Customer{}).Preload("Accounts").Preload("Addresses").First(customer, id); res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			fmt.Println("No customer found with the ID: ", id)

			return nil, fmt.Errorf("no customer found with the ID: %v", id)
		}

		fmt.Println("Error while retrieving a customer with ID: ", id, ", err: ", err)

		return nil, err
	}

	return customer, nil
}

func (c *customer) Update(id int, customerRequest *models.Customer) (updatedCustomer *models.Customer, err error) {
	customerRequest.ID = uint(id)
	if res := c.db.Save(customerRequest); res.Error != nil {
		fmt.Println("Error while updating the customer with ID: ", id, ", err: ", err)

		return nil, err
	}

	if updatedCustomer, err = c.Get(id); err != nil {
		return nil, err
	}

	return updatedCustomer, nil
}

func (c *customer) Delete(id int) (err error) {
	if res := c.db.Delete(&models.Customer{}, id); res.Error != nil {
		fmt.Println("Error while deleting the customer with ID: ", id, ", err: ", err)

		return err
	}

	return nil
}
