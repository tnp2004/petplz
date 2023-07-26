package service

func (s accountService) Login(email, password string) error {
	err := s.accountRepo.Login(email, password)
	if err != nil {
		return err
	}

	return nil
}
