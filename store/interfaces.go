package store

import "github.com/maha-1030/go-rest-api/models"

type Customer interface {
	GetAll() (customers []models.Customer, err error)
	Create(customer *models.Customer) (newCustomer *models.Customer, err error)
	Get(id int) (customer *models.Customer, err error)
	Update(id int, customer *models.Customer) (updatedCustomer *models.Customer, err error)
	Delete(id int) (err error)
}
