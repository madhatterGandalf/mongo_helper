package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var CitiesColl *CitiesCollection
var CountriesColl *CountriesCollection
var PhoneColl *PhoneCollection
var TownColl *TownCollection

func init() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://192.168.243.131:27017"))
	if err != nil {
		panic(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	CitiesColl = &CitiesCollection{collection: client.Database("world").Collection("cities"), ctx: ctx}
	CountriesColl = &CountriesCollection{collection: client.Database("country").Collection("countries"), ctx: ctx}
	PhoneColl = &PhoneCollection{collection: client.Database("phone").Collection("phones"), ctx: ctx}
	TownColl = &TownCollection{collection: client.Database("town").Collection("towns"), ctx: ctx}
}
