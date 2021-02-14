package seeders

import (
	"encoding/json"
	"log"
	"resident-graphql/models/rajaongkir"
	"resident-graphql/models/tables"
	"resident-graphql/services"

	"gorm.io/gorm"
)

func seedProvince(db *gorm.DB) {
	resBody := services.FetchFromRajaongkir("/province")
	var (
		response rajaongkir.Prov
		prov     tables.Provinces
	)

	if err := json.Unmarshal(resBody, &response); err != nil {
		log.Fatalln("Error -> ", err.Error())
	}

	for i := 0; i < len(response.RajaOngkir.ProvinceResults); i++ {
		prov.ProvinceName = response.RajaOngkir.ProvinceResults[i].Province
		prov.ID = 0
		db.Create(&prov)
	}
	log.Println("Seed Province created")
}
