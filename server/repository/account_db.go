package repository

import (
	"context"
	"encoding/json"
	"fmt"

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

func (r accountRepositoryDB) Create(account Account) (*Account, error) {
	return nil, nil
}
func (r accountRepositoryDB) GetAccount(id string) (*Account, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)

	coll := r.db.Database("Renter").Collection("Account")
	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{Key: "_id", Value: objectId}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", id)
		return nil, err
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)

	return nil, nil
}
