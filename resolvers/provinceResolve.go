package resolvers

import (
	"math/rand"
	"resident-graphql/connection"
	"resident-graphql/models/tables"
	"time"

	"github.com/graphql-go/graphql"
)

// CreateProvinceResolver func
func CreateProvinceResolver(p graphql.ResolveParams) (interface{}, error) {
	db := *connection.GetConnection()
	rand.Seed(time.Now().UnixNano())

	var province tables.Provinces

	province.ID = uint(rand.Intn(100000))
	province.Name = p.Args["name"].(string)

	db.Create(&province)

	return province, nil
}

// ReadProvinceResolver func
func ReadProvinceResolver(p graphql.ResolveParams) (interface{}, error) {
	db := *connection.GetConnection()
	var province []tables.Provinces
	db.Preload("District").Find(&province)
	return province, nil
}

// UpdateProvinceResolver func
func UpdateProvinceResolver(p graphql.ResolveParams) (interface{}, error) {
	db := *connection.GetConnection()
	id, _ := p.Args["id"].(int)
	name, _ := p.Args["name"].(string)

	var province tables.Provinces

	db.Model(&province).Where("id = ?", id).Update("name", name)

	return province, nil
}

// DeleteProvinceResolver func
func DeleteProvinceResolver(p graphql.ResolveParams) (interface{}, error) {
	db := *connection.GetConnection()
	id, _ := p.Args["id"].(int)

	var province tables.Provinces

	db.Delete(&province, id)

	return province, nil
}
