package address

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/maha-1030/go-rest-api/models"
	"github.com/maha-1030/go-rest-api/store"
)

type address struct {
	db *gorm.DB
}

func NewAddress(db *gorm.DB) store.Address {
	return &address{
		db: db,
	}
}

func (a *address) GetAll() (addresses []models.Address, err error) {
	res := a.db.Find(&addresses)
	if res.Error != nil {
		fmt.Println("Error while retrieving addresses data, err: ", res.Error)

		return nil, res.Error
	}

	return addresses, nil
}

func (a *address) Create(address *models.Address) (newAddress *models.Address, err error) {
	if res := a.db.Create(address); res.Error != nil {
		fmt.Println("Error while creating the new address, err: ", err)

		return nil, err
	}

	return address, nil
}

func (a *address) Get(id int) (address *models.Address, err error) {
	address = &models.Address{}

	if res := a.db.First(address, id); res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			fmt.Println("No address found with the ID: ", id)

			return nil, fmt.Errorf("no address found with the ID: %v", id)
		}

		fmt.Println("Error while retrieving a address with ID: ", id, ", err: ", err)

		return nil, err
	}

	return address, nil
}

func (a *address) Update(id int, addressRequest *models.Address) (updatedAddress *models.Address, err error) {
	addressRequest.ID = uint(id)
	if res := a.db.Save(addressRequest); res.Error != nil {
		fmt.Println("Error while updating the address with ID: ", id, ", err: ", err)

		return nil, err
	}

	if updatedAddress, err = a.Get(id); err != nil {
		return nil, err
	}

	return updatedAddress, nil
}

func (a *address) Delete(id int) (err error) {
	if res := a.db.Delete(&models.Address{}, id); res.Error != nil {
		fmt.Println("Error while deleting the address with ID: ", id, ", err: ", err)

		return err
	}

	return nil
}
