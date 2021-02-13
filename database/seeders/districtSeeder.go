package seeders

import (
	"encoding/json"
	"log"
	"resident-graphql/models/rajaongkir"
	"resident-graphql/models/tables"
	"resident-graphql/services"
	"strconv"

	"gorm.io/gorm"
)

func seedDistrict(db *gorm.DB) {
	resBody := services.FetchFromRajaongkir("/city")
	var (
		response rajaongkir.City
		dist     tables.Districts
	)

	if err := json.Unmarshal(resBody, &response); err != nil {
		log.Fatalln("Error -> ", err.Error())
	}

	for i := 0; i < 100; i++ {
		uProvID, _ := strconv.ParseUint(response.RajaOngkir.CityResults[i].ProvinceID, 10, 32)
		dist.Name = response.RajaOngkir.CityResults[i].CityName
		dist.ProvinceID = uint(uProvID)
		dist.ID = 0
		db.Create(&dist)
	}
	log.Println("Seed District created")
}
