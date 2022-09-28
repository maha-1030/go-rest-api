package service

import "github.com/maha-1030/go-rest-api/models"

type Customer interface {
	GetAll() (customers []models.Customer, err error)
	Create(customer *models.Customer) (newCustomer *models.Customer, err error)
	Get(id string) (customer *models.Customer, err error)
	Update(id string, customer *models.Customer) (updatedCustomer *models.Customer, err error)
	Delete(id string) (err error)
}

type Address interface {
	GetAll() (addresses []models.Address, err error)
	Create(customerID string, addresses *models.Address) (newAddres *models.Address, err error)
	Get(id, customerID string) (address *models.Address, err error)
	Update(id, customerID string, address *models.Address) (updatedAddress *models.Address, err error)
	Delete(id, customerID string) (err error)
}

type Account interface {
	GetAll() (accounts []models.Account, err error)
	Create(customerID string, account *models.Account) (newAccount *models.Account, err error)
	Get(id, customerID string) (account *models.Account, err error)
	UpdateBalance(id, customerID string, account *models.Account) (updatedAccount *models.Account, err error)
	Delete(id, customerID string) (err error)
}
