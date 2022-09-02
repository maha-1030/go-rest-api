package account

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/maha-1030/go-rest-api/models"
	"github.com/maha-1030/go-rest-api/store"
)

type account struct {
	db *gorm.DB
}

func NewAccount(db *gorm.DB) store.Account {
	return &account{
		db: db,
	}
}

func (a *account) GetAll() (accounts []models.Account, err error) {
	res := a.db.Find(&accounts)
	if res.Error != nil {
		fmt.Println("Error while retrieving accounts data, err: ", res.Error)

		return nil, res.Error
	}

	return accounts, nil
}

func (a *account) Create(account *models.Account) (newAccount *models.Account, err error) {
	if res := a.db.Create(account); res.Error != nil {
		fmt.Println("Error while creating the new account, err: ", err)

		return nil, err
	}

	return account, nil
}

func (a *account) Get(id int) (account *models.Account, err error) {
	account = &models.Account{}

	if res := a.db.First(account, id); res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			fmt.Println("No account found with the ID: ", id)

			return nil, fmt.Errorf("no account found with the ID: %v", id)
		}

		fmt.Println("Error while retrieving a account with ID: ", id, ", err: ", err)

		return nil, err
	}

	return account, nil
}

func (a *account) Update(id int, accountRequest *models.Account) (updatedAccount *models.Account, err error) {
	accountRequest.ID = uint(id)
	if res := a.db.Save(accountRequest); res.Error != nil {
		fmt.Println("Error while updating the account with ID: ", id, ", err: ", err)

		return nil, err
	}

	if updatedAccount, err = a.Get(id); err != nil {
		return nil, err
	}

	return updatedAccount, nil
}

func (a *account) Delete(id int) (err error) {
	if res := a.db.Delete(&models.Account{}, id); res.Error != nil {
		fmt.Println("Error while deleting the account with ID: ", id, ", err: ", err)

		return err
	}

	return nil
}
