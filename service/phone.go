package service

import (
	"fmt"
	"github.com/dickymrlh/mongo_helper/dao"
	"math/rand"
)

type PhoneService int

func (p *PhoneService) PlayAroundWithPhone() {
	p.populatePhones(900, 6150000, 6250000)
}

func (*PhoneService) populatePhones(area, start, stop int) {
	i := start
	for i < stop {
		country := rand.Intn(8)
		num := int64((country * 1e10) + (area * 1e7) + i)
		fullNumber := fmt.Sprintf("+%d %d-%d", country, area, i)
		err := dao.PhoneColl.InsertOne(dao.Phone{
			ID: num,
			Components: dao.Component{
				Country: country,
				Area:    area,
				Prefix:  (i / 1e4),
				Number:  i,
			},
			Display: fullNumber,
		})
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			i++
			fmt.Printf("inserted number: %s\n", fullNumber)
		}
	}

	fmt.Println("DONE")
}
