package service

func (s accountService) Login(email, password string) (*AccountResponse, error) {
	account, err := s.accountRepo.Login(email, password)
	if err != nil {
		return nil, err
	}

	accountRes := AccountResponse{
		AccountID: account.AccountID,
		Username:  account.Username,
		Email:     account.Email,
		Name:      account.Name,
		Gender:    account.Gender,
		Age:       account.Age,
		Money:     account.Money,
		Image_uri: account.Image_uri,
	}

	return &accountRes, nil
}
