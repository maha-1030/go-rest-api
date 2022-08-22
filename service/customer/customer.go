package customer

import (
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
func (c *customer) Get() (customers []models.Customer, err error) {
	return c.cs.Get()
}
