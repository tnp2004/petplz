package service

import "go.mongodb.org/mongo-driver/bson/primitive"

type NewAccount struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Age       int    `json:"age"`
	Money     int    `json:"money"`
	Image_uri string `json:"image_uri"`
}

type AccountResponse struct {
	AccountID primitive.ObjectID `json:"accountId" bson:"_id"`
	Email     string             `json:"email"`
	Name      string             `json:"name"`
	Gender    string             `json:"gender"`
	Age       int                `json:"age"`
	Money     int                `json:"money"`
	Image_uri string             `json:"image_uri"`
}

type AccountService interface {
	NewAccount(NewAccount) error
	GetUserAccount(id string) (*AccountResponse, error)
}
