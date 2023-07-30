package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/tnp2004/petplz/handler"
	"github.com/tnp2004/petplz/repository"
	"github.com/tnp2004/petplz/service"
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

	accountRepositoryDB := repository.NewAccountRepositoryDB(dbClient)
	accountService := service.NewAccountService(accountRepositoryDB)
	accountHandler := handler.NewAccoutHandler(accountService)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Get("/", accountHandler.Greeting)

	app.Get("/:id", accountHandler.GetAccount)
	app.Post("/create", accountHandler.CreateAccount)

	app.Post("/login", accountHandler.Login)
	app.Post("/logout", accountHandler.Logout)
	app.Post("/test/greet", accountHandler.LoginAuth, accountHandler.Greeting)

	app.Listen(":3000")
}

func initDatabase(uri string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer cancel()

	return client
}

func loadDotEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}
