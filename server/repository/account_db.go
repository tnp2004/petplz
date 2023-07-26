package repository

import (
	"context"
	"encoding/json"
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

	coll := r.db.Database("Renter").Collection("Account")
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
		fmt.Println(err)
	}

	coll := r.db.Database("Renter").Collection("Account")
	var result Account
	err = coll.FindOne(context.TODO(), bson.D{{Key: "_id", Value: objectId}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the id %s\n", id)
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

	return account, nil
}
