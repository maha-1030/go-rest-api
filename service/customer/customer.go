package customer

import (
	"fmt"

	"github.com/maha-1030/go-rest-api/models"
	"github.com/maha-1030/go-rest-api/service"
	"github.com/maha-1030/go-rest-api/store"
)

type customer struct {
	cs store.Customer
}

func NewCustomer(cs store.Customer) service.Customer {
	return &customer{
		cs: cs,
	}

}
func (c *customer) GetAll() (customers []models.Customer, err error) {
	return c.cs.GetAll()
}

func (c *customer) Create(customerRequest *models.Customer) (newCustomer *models.Customer, err error) {
	return c.cs.Create(customerRequest)
}

func (c *customer) Get(idString string) (cust *models.Customer, err error) {
	id, err := service.GetID(idString)
	if err != nil {
		fmt.Println("Got invalid ID in the Get Customer request")

		return nil, err
	}

	return c.cs.Get(id)
}

func (c *customer) Update(idString string, customerRequest *models.Customer) (updatedCustomer *models.Customer, err error) {
	id, err := service.GetID(idString)
	if err != nil {
		fmt.Println("Got invalid ID in the Update Customer request")

		return nil, err
	}

	existingCustomer, err := c.cs.Get(id)
	if err != nil {
		fmt.Println("Unable find customer with given ID in Update Customer request, err: ", err)

		return nil, err
	}

	if customerRequest.Name == "" {
		customerRequest.Name = existingCustomer.Name
	}

	if customerRequest.Age == 0 {
		customerRequest.Age = existingCustomer.Age
	}

	if customerRequest.PhoneNumber == "" {
		customerRequest.PhoneNumber = existingCustomer.PhoneNumber
	}

	if customerRequest.Gender == "" {
		customerRequest.Gender = existingCustomer.Gender
	}

	return c.cs.Update(id, customerRequest)
}

func (c *customer) Delete(idString string) (err error) {
	id, err := service.GetID(idString)
	if err != nil {
		fmt.Println("Got invalid ID in the Delete Customer request")

		return err
	}

	_, err = c.cs.Get(id)
	if err != nil {
		fmt.Println("Unable to find customer with given ID in Delete Customer request, err: ", err)

		return err
	}

	return c.cs.Delete(id)
}
