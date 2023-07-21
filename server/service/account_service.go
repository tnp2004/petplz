package service

import (
	"github.com/tnp2004/Renter/repository"
)

type accountService struct {
	accountRepo repository.AccountRepository
}

func NewAccountService(accountRepo repository.AccountRepository) AccountService {
	return accountService{accountRepo}
}

func (s accountService) NewAccount(account NewAccount) (*AccountResponse, error) {
	return nil, nil
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
