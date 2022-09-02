package store

import "github.com/maha-1030/go-rest-api/models"

type Customer interface {
	GetAll() (customers []models.Customer, err error)
	Create(customer *models.Customer) (newCustomer *models.Customer, err error)
	Get(id int) (customer *models.Customer, err error)
	Update(id int, customer *models.Customer) (updatedCustomer *models.Customer, err error)
	Delete(id int) (err error)
}

type Address interface {
	GetAll() (addresses []models.Address, err error)
	Create(addresses *models.Address) (newAddres *models.Address, err error)
	Get(id int) (address *models.Address, err error)
	Update(id int, address *models.Address) (updatedAddress *models.Address, err error)
	Delete(id int) (err error)
}

type Account interface {
	GetAll() (accounts []models.Account, err error)
	Create(account *models.Account) (newAccount *models.Account, err error)
	Get(id int) (account *models.Account, err error)
	Update(id int, account *models.Account) (updatedAccount *models.Account, err error)
	Delete(id int) (err error)
}
