package services

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// FetchFromRajaongkir from api.rajaongkir.com
func FetchFromRajaongkir(data string) []byte {
	baseURL := "https://api.rajaongkir.com/starter"
	apiKey := os.Getenv("RAJAONGKIR_APIKEY")
	res, err := http.Get(baseURL + data + "?key=" + apiKey)
	if err != nil {
		log.Fatalln("Error -> ", err.Error())
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("Error -> ", err.Error())
	}

	return resBody
}
