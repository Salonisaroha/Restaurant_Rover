package main

import (
	"flag"

	//"github.com/Salonisaroha/api"

	"github.com/gofiber/fiber/v2"
)

func handleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "james foo"})
}
func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "this is working fyn"})
}

func main() {
	// listenAddar := flag.String("listenAddr", ":5000", "The listen address of the API server ")
	// flag.Parse()
	// app := fiber.New()
	// app.Get("/user", api.HandleGetUsers)
	// app.Get("/user/:id", api.HandleGetUser)

	// app.Listen(*listenAddar)
	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	flag.Parse()
	app := fiber.New()

	appv1 := app.Group("/api/v1")
	app.Get("/foo", handleFoo)
	appv1.Get("/user", handleUser)
	app.Listen(*listenAddr)
}
