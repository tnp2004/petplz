package service

import "go.mongodb.org/mongo-driver/bson/primitive"

type NewAccount struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Gender    string `json:"gender"`
	Age       int    `json:"age"`
	Image_uri string `json:"image_uri"`
}

type AccountResponse struct {
	AccountID primitive.ObjectID `json:"accountId" bson:"_id"`
	Username  string             `json:"username"`
	Email     string             `json:"email"`
	Gender    string             `json:"gender"`
	Age       int                `json:"age"`
	Money     int                `json:"money"`
	Image_uri string             `json:"image_uri"`
}

type LoginRequired struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AccountService interface {
	Register(NewAccount) error
	Login(email, password string) (*AccountResponse, error)
	GetUserAccount(id string) (*AccountResponse, error)
}
