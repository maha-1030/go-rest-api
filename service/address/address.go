package address

import (
	"fmt"

	"github.com/maha-1030/go-rest-api/models"
	"github.com/maha-1030/go-rest-api/service"
	"github.com/maha-1030/go-rest-api/store"
)

type addres struct {
	as store.Address
	cs store.Customer
}

func NewAddress(as store.Address, cs store.Customer) service.Address {
	return &addres{
		as: as,
		cs: cs,
	}
}

func (a *addres) GetAll() (addresses []models.Address, err error) {
	return a.as.GetAll()
}

func (a *addres) Create(customerIDString string, addressRequest *models.Address) (newAddress *models.Address, err error) {
	customerID, err := service.GetID(customerIDString)
	if err != nil {
		fmt.Println("Got invalid customerID in the Create Address request")

		return nil, err
	}

	_, err = a.cs.Get(customerID)
	if err != nil {
		fmt.Println("Unable to find the customer in Create Address request, customerID: ", customerID)

		return nil, err
	}

	addressRequest.CustomerID = uint(customerID)

	return a.as.Create(addressRequest)
}

func (a *addres) Get(idString, customerIDString string) (addr *models.Address, err error) {
	id, err := service.GetID(idString)
	if err != nil {
		fmt.Println("Got invalid accountID in the Get Address request")

		return nil, err
	}

	customerID, err := service.GetID(customerIDString)
	if err != nil {
		fmt.Println("Got invalid customerID in the Get Address request")

		return nil, err
	}

	_, err = a.cs.Get(customerID)
	if err != nil {
		fmt.Println("Unable to find the customer in Get Address request, customerID: ", customerID)

		return nil, err
	}

	existingAddress, err := a.as.Get(id)
	if err != nil {
		fmt.Println("Unable to find the address in Get Address request, accountID: ", customerID)

		return nil, err
	}

	if existingAddress.CustomerID != uint(customerID) {
		return nil, fmt.Errorf("address with ID: %v doesn't belongs to the customer with ID: %v", id, customerID)
	}

	return existingAddress, nil
}

func (a *addres) Update(customerIDString string, addressRequest *models.Address) (updatedAddress *models.Address, err error) {
	if addressRequest.ID < 1 {
		fmt.Println("Got invalid AddressID in the Update Address request")

		return nil, fmt.Errorf("invalid address id: %v", addressRequest.ID)
	}

	customerID, err := service.GetID(customerIDString)
	if err != nil {
		fmt.Println("Got invalid customerID in the Update Address request")

		return nil, err
	}

	_, err = a.cs.Get(customerID)
	if err != nil {
		fmt.Println("Unable to find customer with given ID in Update Address request, err: ", err)

		return nil, err
	}

	existingAddress, err := a.as.Get(int(addressRequest.ID))
	if err != nil {
		fmt.Println("Unable to find address with given ID in Update Address request, err: ", err)

		return nil, err
	}

	if existingAddress.CustomerID != uint(customerID) {
		return nil, fmt.Errorf("address with ID: %v doesn't belongs to the customer with ID: %v", addressRequest.ID, customerID)
	}

	if addressRequest.Line1 != "" {
		existingAddress.Line1 = addressRequest.Line1
	}

	if addressRequest.Line2 != "" {
		existingAddress.Line2 = addressRequest.Line2
	}

	if addressRequest.City != "" {
		existingAddress.City = addressRequest.City
	}

	if addressRequest.State != "" {
		existingAddress.State = addressRequest.State
	}

	if addressRequest.PinCode != "" {
		existingAddress.PinCode = addressRequest.PinCode
	}

	return a.as.Update(int(addressRequest.ID), existingAddress)
}

func (a *addres) Delete(idString, customerIDString string) (err error) {
	id, err := service.GetID(idString)
	if err != nil {
		fmt.Println("Got invalid addressID in the Delete Address request")

		return err
	}

	customerID, err := service.GetID(customerIDString)
	if err != nil {
		fmt.Println("Got invalid customerID in the Delete Address request")

		return err
	}

	_, err = a.cs.Get(customerID)
	if err != nil {
		fmt.Println("Unable to find customer with given ID in Delete Address request, err: ", err)

		return err
	}

	existingAddress, err := a.as.Get(id)
	if err != nil {
		fmt.Println("Unable to find address with given ID in Delete Address request, err: ", err)

		return err
	}

	if existingAddress.CustomerID != uint(customerID) {
		return fmt.Errorf("address with ID: %v doesn't belongs to the customer with ID: %v", idString, customerIDString)
	}

	return a.as.Delete(id)
}
