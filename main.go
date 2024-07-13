package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/Salonisaroha/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func handleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "james foo"})
}
func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "this is working fyn"})
}

const dburi = "mongodb://localhost:27017"
const dbname = "Hotel-reservation"
const userColl = "users"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	coll := client.Database(dbname).Collection(userColl)

	user := types.User{
		FirstName: "James",
		LastName:  "At the water cooler",
	}
	res, err := coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	flag.Parse()
	app := fiber.New()

	appv1 := app.Group("/api/v1")
	app.Get("/foo", handleFoo)
	appv1.Get("/user", handleUser)
	app.Listen(*listenAddr)
}
