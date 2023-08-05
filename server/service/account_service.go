package service

import (
	"time"

	"github.com/tnp2004/petplz/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type accountService struct {
	accountRepo repository.AccountRepository
}

func NewAccountService(accountRepo repository.AccountRepository) AccountService {
	return accountService{accountRepo}
}

func (s accountService) Register(accountRegister NewAccount) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(accountRegister.Password), 12)
	if err != nil {
		return err
	}

	accountReg := repository.Account{
		AccountID:  primitive.NewObjectID(),
		Username:   accountRegister.Username,
		Email:      accountRegister.Email,
		Password:   string(hashedPassword),
		Gender:     accountRegister.Gender,
		Age:        accountRegister.Age,
		Money:      0,
		Image_uri:  "https://static.printler.com/cache/0/8/1/1/8/c/08118cb095d702b52289a030f9ba1188e345c33b.jpg",
		Created_at: time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}

	err = s.accountRepo.Register(accountReg)
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
		Username:  account.Username,
		Email:     account.Email,
		Gender:    account.Gender,
		Age:       account.Age,
		Money:     account.Money,
		Image_uri: account.Image_uri,
	}

	return &accountRes, nil
}
