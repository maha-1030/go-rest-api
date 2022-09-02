package account

import (
	"fmt"

	"github.com/maha-1030/go-rest-api/models"
	"github.com/maha-1030/go-rest-api/service"
	"github.com/maha-1030/go-rest-api/store"
)

type account struct {
	as store.Account
	cs store.Customer
}

func NewAccount(as store.Account, cs store.Customer) service.Account {
	return &account{
		as: as,
		cs: cs,
	}

}
func (a *account) GetAll() (accounts []models.Account, err error) {
	return a.as.GetAll()
}

func (a *account) Create(customerIDString string, accountRequest *models.Account) (newAccount *models.Account, err error) {
	customerID, err := service.GetID(customerIDString)
	if err != nil {
		fmt.Println("Got invalid customerID in the Create Account request")

		return nil, err
	}

	_, err = a.cs.Get(customerID)
	if err != nil {
		fmt.Println("Unable to find the customer in Create Account request, customerID: ", customerID)

		return nil, err
	}

	accountRequest.CustomerID = uint(customerID)

	return a.as.Create(accountRequest)
}

func (a *account) Get(idString, customerIDString string) (acc *models.Account, err error) {
	id, err := service.GetID(idString)
	if err != nil {
		fmt.Println("Got invalid accountID in the Get Account request")

		return nil, err
	}

	customerID, err := service.GetID(customerIDString)
	if err != nil {
		fmt.Println("Got invalid customerID in the Get Account request")

		return nil, err
	}

	_, err = a.cs.Get(customerID)
	if err != nil {
		fmt.Println("Unable to find the customer in Get Account request, customerID: ", customerID)

		return nil, err
	}

	existingAccount, err := a.as.Get(id)
	if err != nil {
		fmt.Println("Unable to find the account in Get Account request, accountID: ", customerID)

		return nil, err
	}

	if existingAccount.CustomerID != uint(customerID) {
		return nil, fmt.Errorf("account with ID: %v doesn't belongs to the customer with ID: %v", id, customerID)
	}

	return existingAccount, nil
}

func (a *account) UpdateBalance(customerIDString string, accountRequest *models.Account) (updatedAccount *models.Account, err error) {
	if accountRequest.ID < 1 {
		fmt.Println("Got invalid AccountID in the Update Account Balance request")

		return nil, fmt.Errorf("invalid account id: %v", accountRequest.ID)
	}

	customerID, err := service.GetID(customerIDString)
	if err != nil {
		fmt.Println("Got invalid customerID in the Update Account Balance request")

		return nil, err
	}

	_, err = a.cs.Get(customerID)
	if err != nil {
		fmt.Println("Unable to find customer with given ID in Update Account Balance request, err: ", err)

		return nil, err
	}

	existingAccount, err := a.as.Get(int(accountRequest.ID))
	if err != nil {
		fmt.Println("Unable to find account with given ID in Update Account Balance request, err: ", err)

		return nil, err
	}

	if existingAccount.CustomerID != uint(customerID) {
		return nil, fmt.Errorf("account with ID: %v doesn't belongs to the customer with ID: %v", accountRequest.ID, customerID)
	}

	existingAccount.Balance = accountRequest.Balance

	return a.as.Update(int(accountRequest.ID), existingAccount)
}

func (a *account) Delete(idString, customerIDString string) (err error) {
	id, err := service.GetID(idString)
	if err != nil {
		fmt.Println("Got invalid accountID in the Delete Account request")

		return err
	}

	customerID, err := service.GetID(customerIDString)
	if err != nil {
		fmt.Println("Got invalid customerID in the Delete Account request")

		return err
	}

	_, err = a.cs.Get(customerID)
	if err != nil {
		fmt.Println("Unable to find customer with given ID in Delete Account request, err: ", err)

		return err
	}

	existingAccount, err := a.as.Get(id)
	if err != nil {
		fmt.Println("Unable to find account with given ID in Delete Account request, err: ", err)

		return err
	}

	if existingAccount.CustomerID != uint(customerID) {
		return fmt.Errorf("account with ID: %v doesn't belongs to the customer with ID: %v", idString, customerIDString)
	}

	return a.as.Delete(id)
}
