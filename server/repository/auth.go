package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func (r accountRepositoryDB) Login(email, password string) error {

	coll := r.db.Database("Renter").Collection("Account")
	var result Account
	err := coll.FindOne(context.TODO(), bson.D{{Key: "email", Value: email}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the email %s\n", email)
		return err
	}
	if err != nil {
		return err
	}

	accByte, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return err
	}

	var account *Account
	err = json.Unmarshal(accByte, &account)
	if err != nil {
		return err
	}

	// invalid password
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil {
		return err
	}

	// valid password
	return nil
}
