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

func (r accountRepositoryDB) Create(account Account) (*Account, error) {

	coll := r.db.Database("Renter").Collection("Account")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	account.AccountID = primitive.NewObjectID()
	account.Created_at = time.Now().Format("2006-01-02T15:04:05Z07:00")

	result, err := coll.InsertOne(ctx, account)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result.InsertedID)

	return nil, nil
}
func (r accountRepositoryDB) GetAccount(id string) (*Account, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
	}

	coll := r.db.Database("Renter").Collection("Account")
	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{Key: "_id", Value: objectId}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the id %s\n", id)
		return nil, err
	}
	if err != nil {
		fmt.Println("hi")
		panic(err)
	}
	accountData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", accountData)

	return nil, nil
}
