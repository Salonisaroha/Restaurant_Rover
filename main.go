package main

import (
	"flag"

	"github.com/Salonisaroha/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddar := flag.String("listenAddr", ":5000", "The listen address of the API server ")
	flag.Parse()
	app := fiber.New()
	app.Get("/user", api.HandleGetUsers)
	app.Get("/user/:id", api.HandleGetUser)

	app.Listen(*listenAddar)
}
