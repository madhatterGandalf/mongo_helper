package service

import (
	"fmt"
	"github.com/dickymrlh/mongo_helper/dao"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CountryService int

func (c *CountryService) PlayAroundWithCountry() {
	fmt.Println("###################################################################################")
	err := c.DelTown()
	if err != nil {
		panic(err)
	}

	err = c.InsertTown()
	if err != nil {
		panic(err)
	}

	// find With elemMatch
	// It specifies that if a document (or nested document)
	// matches all of our criteria
	countries, err := dao.CountriesColl.Find(bson.M{
		"exports.foods": bson.M{
			"$elemMatch": bson.M{
				"name":  "bacon",
				"tasty": true,
			},
		},
	}, options.Find())
	if err != nil {
		panic(err)
	}
	fmt.Println(countries)
	fmt.Println()
	fmt.Println("###################################################################################")

	// find With or
	countries, err = dao.CountriesColl.Find(bson.M{
		"$or": []bson.M{
			bson.M{"_id": "mx"},
			bson.M{"name": "United States"},
		},
	}, options.Find().SetProjection(bson.D{{"_id", 1}}))
	if err != nil {
		panic(err)
	}
	fmt.Println(countries)
	fmt.Println()
	fmt.Println("###################################################################################")

	// remove with bad bacon
	removedCount, err := dao.CountriesColl.Remove(bson.M{
		"exports.foods": bson.M{
			"$elemMatch": bson.M{
				"name":  "bacon",
				"tasty": false,
			},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(removedCount)
	fmt.Println()
	fmt.Println("###################################################################################")
}

func (c *CountryService) InsertTown() error {

	countries := []dao.Country{
		dao.Country{
			ID:   "us",
			Name: "United States",
			Exports: dao.Export{
				[]dao.Food{
					dao.Food{Name: "bacon", Tasty: true},
					dao.Food{Name: "burgers"},
				},
			},
		},
		dao.Country{
			ID:   "ca",
			Name: "Canada",
			Exports: dao.Export{
				[]dao.Food{
					dao.Food{Name: "bacon", Tasty: false},
					dao.Food{Name: "syrup", Tasty: true},
				},
			},
		},
		dao.Country{
			ID:   "mx",
			Name: "Mexico",
			Exports: dao.Export{
				[]dao.Food{
					dao.Food{Name: "salsa", Tasty: true, Condiment: true},
				},
			},
		},
	}

	err := dao.CountriesColl.InsertMany(countries)
	if err != nil {
		return err
	}

	return nil
}

func (c *CountryService) DelTown() error {
	_, err := dao.CountriesColl.Remove(bson.M{
		"exports.foods": bson.M{
			"$elemMatch": bson.M{
				"name": bson.M{
					"$ne": "123",
				},
			},
		},
	})
	if err != nil {
		return err
	}

	return nil
}
