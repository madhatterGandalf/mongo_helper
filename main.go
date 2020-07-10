package main

import (
	_ "github.com/dickymrlh/mongo_helper/dao"
	"github.com/dickymrlh/mongo_helper/service"
)

func main() {
	countryService := new(service.CountryService)
	countryService.PlayAroundWithCountry()

}
