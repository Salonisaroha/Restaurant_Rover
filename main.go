package main

import (
	"context"
	"flag"
	"log"

	"github.com/Salonisaroha/api"
	"github.com/Salonisaroha/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"
const dbname = "Hotel-reservation"

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		log.Printf("Error: %v\n", err)
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	// Handlers initialization
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	app := fiber.New(config)
	appv1 := app.Group("/api/v1")
	appv1.Put("user/:id", userHandler.HandlePutUser)
    appv1.Delete("/user/:id", userHandler.HandleDeleteUser)
	appv1.Post("/user", userHandler.HandlePostUser)
	appv1.Get("/users", userHandler.HandleGetUsers)   // Fetch all users
	appv1.Get("/user/:id", userHandler.HandleGetUser) // Fetch single user by ID

	// Start the server
	if err := app.Listen(*listenAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
