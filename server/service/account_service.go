package service

import (
	"time"

	"github.com/tnp2004/petplz/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type accountService struct {
	accountRepo repository.AccountRepository
}

func NewAccountService(accountRepo repository.AccountRepository) AccountService {
	return accountService{accountRepo}
}

func (s accountService) NewAccount(accountRegister NewAccount) error {
	accountReg := repository.Account{
		AccountID:  primitive.NewObjectID(),
		Email:      accountRegister.Email,
		Password:   accountRegister.Password,
		Name:       accountRegister.Name,
		Gender:     accountRegister.Gender,
		Age:        accountRegister.Age,
		Money:      accountRegister.Money,
		Image_uri:  accountRegister.Image_uri,
		Created_at: time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}

	err := s.accountRepo.Create(accountReg)
	if err != nil {
		return err
	}

	return nil
}
func (s accountService) GetUserAccount(id string) (*AccountResponse, error) {
	account, err := s.accountRepo.GetAccount(id)
	if err != nil {
		return nil, err
	}

	accountRes := AccountResponse{
		AccountID: account.AccountID,
		Email:     account.Email,
		Name:      account.Name,
		Gender:    account.Gender,
		Age:       account.Age,
		Money:     account.Money,
		Image_uri: account.Image_uri,
	}

	return &accountRes, nil
}
