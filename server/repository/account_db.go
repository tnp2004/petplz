package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type accountRepositoryDB struct {
	db *mongo.Client
}

func NewAccountRepositoryDB(db *mongo.Client) AccountRepository {
	return accountRepositoryDB{db}
}

func (r accountRepositoryDB) Register(account Account) error {

	coll := r.db.Database("Petplz").Collection("Account")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	account.AccountID = primitive.NewObjectID()
	account.Created_at = time.Now().Format("2006-01-02T15:04:05Z07:00")

	_, err := coll.InsertOne(ctx, account)
	if err != nil {
		return err
	}

	return nil
}

func (r accountRepositoryDB) GetAccount(id string) (*Account, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("id:%s invalid format", id)
	}

	coll := r.db.Database("Petplz").Collection("Account")
	var result Account
	err = coll.FindOne(context.TODO(), bson.D{{Key: "_id", Value: objectId}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("account id:%s doesn't exist", id)
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

	return account, nil
}

func (r accountRepositoryDB) EmailVerification(email string) error {
	coll := r.db.Database("Petplz").Collection("Account")
	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{Key: "email", Value: email}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		// email is unused
		return nil
	}
	if err != nil {
		return err
	}

	return errors.New("something went wrong")
}

func (r accountRepositoryDB) UsernameVerification(username string) error {
	coll := r.db.Database("Petplz").Collection("Account")
	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{Key: "username", Value: username}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		// username is unused
		return nil
	}
	if err != nil {
		return err
	}

	return errors.New("something went wrong")
}
