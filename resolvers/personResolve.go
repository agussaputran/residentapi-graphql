package resolvers

import (
	"fmt"
	"math/rand"
	"resident-graphql/connection"
	"resident-graphql/models/responses"
	"resident-graphql/models/tables"
	"time"

	"github.com/graphql-go/graphql"
)

// CreatePersonResolver func
func CreatePersonResolver(p graphql.ResolveParams) (interface{}, error) {
	db := *connection.GetConnection()
	rand.Seed(time.Now().UnixNano())

	var person tables.Persons

	person.ID = uint(rand.Intn(100000))
	person.Nip = p.Args["nip"].(string)
	person.FullName = p.Args["full_name"].(string)
	person.FirstName = p.Args["first_name"].(string)
	person.LastName = p.Args["last_name"].(string)
	person.BirthDate = p.Args["birth_date"].(string)
	person.BirthPlace = p.Args["birth_place"].(string)
	person.Gender = p.Args["gender"].(string)
	person.ZoneLocation = p.Args["zone_location"].(string)
	person.SubDistrictID = uint(p.Args["sub_district_id"].(int))

	db.Create(&person)

	return person, nil
}

// ReadPersonResolver func
func ReadPersonResolver(p graphql.ResolveParams) (interface{}, error) {
	db := *connection.GetConnection()
	var (
		person   []tables.Persons
		response []responses.PersonResponse
	)
	db.Model(&person).
		Select(`persons.id, persons.nip, persons.full_name, persons.first_name,
		persons.last_name, persons.birth_date, persons.birth_place, persons.photo_profile_url,
		persons.gender, persons.zone_location, sub_districts.sub_district_name`).
		Joins("left join sub_districts on sub_districts.id = persons.sub_district_id").Scan(&response)
	return response, nil
}

// UpdatePersonResolver func
func UpdatePersonResolver(p graphql.ResolveParams) (interface{}, error) {
	db := *connection.GetConnection()
	id, _ := p.Args["id"].(int)
	nip, _ := p.Args["nip"].(string)
	fullName, _ := p.Args["full_name"].(string)
	firstName, _ := p.Args["first_name"].(string)
	lastName, _ := p.Args["last_name"].(string)
	birthDate, _ := p.Args["birth_date"].(string)
	birthPlace, _ := p.Args["birth_place"].(string)
	gender, _ := p.Args["gender"].(string)
	zoneLoc, _ := p.Args["zone_location"].(string)
	subDistrictID, _ := p.Args["sub_district_id"].(int)

	var person tables.Persons

	db.Model(&person).Where("id = ?", id).Updates(tables.Persons{
		Nip:           nip,
		FullName:      fullName,
		FirstName:     firstName,
		LastName:      lastName,
		BirthDate:     birthDate,
		BirthPlace:    birthPlace,
		Gender:        gender,
		ZoneLocation:  zoneLoc,
		SubDistrictID: uint(subDistrictID),
	})

	return person, nil
}

// DeletePersonResolver func
func DeletePersonResolver(p graphql.ResolveParams) (interface{}, error) {
	db := *connection.GetConnection()
	id, _ := p.Args["id"].(int)

	var person tables.Persons

	db.Delete(&person, id)
	fmt.Println(person)

	return person, nil
}
