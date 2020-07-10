package service

import (
	"fmt"
	"github.com/dickymrlh/mongo_helper/dao"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TownService int

func (t *TownService) PlayAroundWithTown() {
	fmt.Println("###################################################################################")
	_, err := t.InsertTown()
	if err != nil {
		panic(err)
	}

	opt := options.Find()
	// find all
	towns, err := dao.TownColl.Find(opt)
	if err != nil {
		panic(err)
	}
	fmt.Println(towns)
	fmt.Println("=====================================")

	// find with limit and sort
	opt.SetLimit(2)
	opt.SetSort(bson.D{{"population", 1}})

	towns, err = dao.TownColl.Find(opt)
	if err != nil {
		panic(err)
	}
	fmt.Println(towns)
	fmt.Println("=====================================")

	// find One with object ID
	objID, err := primitive.ObjectIDFromHex("5eb60f2bba0293032d0f96bb")
	if err != nil {
		panic(err)
	}

	findOne := options.FindOne()
	town, err := dao.TownColl.FindOne(bson.D{{"_id", objID}}, findOne)
	if err != nil {
		panic(err)
	}
	fmt.Println(town)

	// find One with object ID and return Name of the city only
	findOne.SetProjection(bson.D{{"name", 1}})
	town, err = dao.TownColl.FindOne(bson.D{{"_id", objID}}, findOne)
	if err != nil {
		panic(err)
	}

	// find One With population in range
	// $lt = less than
	// $gt = greater than
	findOne.SetProjection(bson.D{{"name", 1}, {"population", 1}})
	town, err = dao.TownColl.FindOne(bson.M{
		"population": bson.M{
			"$lt": 1000000,
			"$gt": 10000,
		},
	}, findOne)
	if err != nil {
		panic(err)
	}

	// find One matching partial values using regEx
	// options "i" = for case incensitive
	regex := bson.M{"$regex": primitive.Regex{Pattern: "moma", Options: "i"}}
	findOne.SetProjection(bson.D{{"name", 1}, {"famousfor", 1}})
	town, err = dao.TownColl.FindOne(bson.M{"famousfor": regex}, findOne)
	if err != nil {
		panic(err)
	}
	fmt.Println(town)
	fmt.Println()

	// update document add state field
	objID, err = primitive.ObjectIDFromHex("5eb60f2bba0293032d0f96bd")
	if err != nil {
		panic(err)
	}

	err = dao.TownColl.UpdateOne(bson.M{"_id": objID}, bson.M{"$set": bson.M{"state": "OR"}})
	if err != nil {
		panic(err)
	}

	// Increase PortLand Population
	err = dao.TownColl.UpdateOne(bson.M{"_id": objID}, bson.M{"$inc": bson.M{"population": 100000}})
	if err != nil {
		panic(err)
	}
	fmt.Println("###################################################################################")
}

func (t *TownService) InsertTown() (ids []string, err error) {

	towns := []dao.Town{
		dao.Town{
			Name:       "New York",
			Population: 22200000,
			LastCensus: time.Date(2016, 7, 1, 0, 0, 0, 0, time.Local),
			FamousFor:  []string{"the MOMA", "food", "Derek Jeter"},
			Mayor:      dao.Politican{Name: "Bill de Blasio", Party: "I"},
		},
		dao.Town{
			Name:       "Punxsutawney",
			Population: 6200,
			LastCensus: time.Date(2016, 1, 31, 0, 0, 0, 0, time.Local),
			FamousFor:  []string{"Punxsutawney Phil"},
			Mayor:      dao.Politican{Name: "Richard Alexander"},
		},
		dao.Town{
			Name:       "Portland",
			Population: 582000,
			LastCensus: time.Date(2016, 9, 20, 0, 0, 0, 0, time.Local),
			FamousFor:  []string{"berr", "food", "Portlandia"},
			Mayor:      dao.Politican{Name: "Ted Wheeler", Party: "D"},
		},
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		ids, err = dao.TownColl.InsertMany(towns)
	}()
	wg.Wait()
	if err != nil {
		return
	}

	return
}
