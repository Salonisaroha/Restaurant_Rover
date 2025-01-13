package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Salonisaroha/db"
	"github.com/Salonisaroha/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	rooms := [] types.Room{
		{
			Type: types.SingleRoomType,
			BasePrice: 99.9,
		},
		{
            Type:types.DeluxRoomType,
			BasePrice: 199.9,
		},
		{
            Type:types.SeaSideRoomType,
			BasePrice: 122.9,
		},
	}
	
	
	InsertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err!= nil{
		log.Fatal(err)
	}
	for _, room := range rooms{
		room.ID = primitive.NewObjectID()
		room.HotelID = InsertedHotel.ID
		insertedRoom, err := roomStore.InsertRoom(ctx, &room)
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println(insertedRoom)
	}
	
	fmt.Println(InsertedHotel)
	
}