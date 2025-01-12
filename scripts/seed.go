package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Salonisaroha/db"
	"github.com/Salonisaroha/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func main(){
	ctx := context.Background()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err!= nil{
		log.Fatal(err)
	}
	hotelStore := db.NewMongoHotelStore(client, db.DBNAME)
	roomStore := db.NewMongoRoomStore(client, db.DBNAME)
	hotel := types.Hotel{
		Name: "Ontario",
		Location : "Noida",
	}
	room := types.Room{
		Type:types.SingleRoomType,
		BasePrice: 99.9,
	}
	_ = room
	InsertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err!= nil{
		log.Fatal(err)
	}
	room.HotelID = InsertedHotel.ID
	insertedRoom, err := roomStore.InsertRoom(room)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(InsertedHotel)
	fmt.Println(insertedRoom)
}