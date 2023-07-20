package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/tnp2004/Renter/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	loadDotEnv()

	dbUri := os.Getenv("MONGODB_URI")
	if dbUri == "" {
		log.Fatal("mongodb uri is empty")
	}

	dbClient := initDatabase(dbUri)
	defer func() {
		if err := dbClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	accountRepositoryDb := repository.NewAccountRepositoryDB(dbClient)
	// accountRepositoryDb.GetAccount("64b92d1282735f7ff4cdb034")
	accountRepositoryDb.Create(repository.Account{
		Email:     "Jamelnwza007@gmail.com",
		Password:  "jamePassword123",
		Name:      "Jame",
		Gender:    "Male",
		Age:       27,
		Money:     15000,
		Image_uri: "catimage",
	})

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hi there!")
	})

	app.Listen(":3000")
}

func initDatabase(uri string) *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return client
}

func loadDotEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}
