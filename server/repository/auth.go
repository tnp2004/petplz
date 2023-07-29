package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func (r accountRepositoryDB) Login(email, password string) (*Account, error) {

	coll := r.db.Database("Renter").Collection("Account")
	var result Account
	err := coll.FindOne(context.TODO(), bson.D{{Key: "email", Value: email}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the email %s\n", email)
		return nil, err
	}
	if err != nil {
		return nil, err

	}

	accByte, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return nil, err
	}

	var account *Account
	err = json.Unmarshal(accByte, &account)
	if err != nil {
		return nil, err
	}

	// invalid password
	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password)); err != nil {
		return nil, err
	}

	// valid password
	return account, nil
}
