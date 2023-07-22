package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	AccountID  primitive.ObjectID `bson:"_id"`
	Email      string
	Password   string
	Name       string
	Gender     string
	Age        int
	Money      int
	Image_uri  string
	Created_at string
}

type AccountRepository interface {
	Create(Account) error
	GetAccount(id string) (*Account, error)
}
