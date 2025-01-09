package api

import (
	"bytes"
	"context"
	"encoding/json"
	
	"log"
	"net/http/httptest"
	"testing"

	"github.com/Salonisaroha/db"
	"github.com/Salonisaroha/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	testmongodburi = "mongodb://localhost:27017"
	dbname         = "hotel-reservation-test"
)

type testdb struct {
	db.UserStore
}

func (tdb *testdb) teardown(t *testing.T) {
	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testmongodburi))
	if err != nil {
		log.Fatal(err)
	}
	return &testdb{
		UserStore: db.NewMongoUserStore(client, dbname),
	}
}

func TestPostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	// Add test logic for posting a user
	app := fiber.New()
	UserHandler := NewUserHandler(tdb.UserStore)
	app.Post("/", UserHandler.HandlePostUser)
	params := types.CreateUserParams{
		Email:     "some@foo.com",
		FirstName: "james",
		LastName:  "Foo",
		Password:  "lkjhgfdsaqwerty",
	}
	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test((req))
	if err != nil {
		t.Error(err)
	}
	var user types.User
	json.NewDecoder(resp.Body).Decode(&user)
	if len(user.ID)==0{
		t.Errorf("expecting a user id to be sent")
	}
	if len(user.EncryptedPassword)>0{
		t.Errorf("expecting the encryptedPassword not to be included in the json response")
	}
	if user.FirstName!= params.FirstName{
		t.Errorf("expected firstname %s but got %s", params.FirstName, user.FirstName)
	}
	if user.LastName!= params.LastName{
		t.Errorf("expected last name %s but got %s", params.LastName, user.LastName)
	}
	if user.Email!= params.Email{
		t.Errorf("expected email %s but got %s", params.Email, user.Email)
	}
}
