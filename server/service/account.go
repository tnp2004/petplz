package service

import "go.mongodb.org/mongo-driver/bson/primitive"

type NewAccount struct {
	Email     string
	Name      string
	Gender    string
	Age       int
	Money     int
	Image_uri string
}

type AccountResponse struct {
	AccountID primitive.ObjectID `bson:"_id"`
	Email     string
	Name      string
	Gender    string
	Age       int
	Money     int
	Image_uri string
}

type AccountService interface {
	NewAccount(NewAccount) (*AccountResponse, error)
	GetUserAccount(id string) (*AccountResponse, error)
}
