package resolvers

import (
	"fmt"
	"math/rand"
	"resident-graphql/connection"
	"resident-graphql/helper"
	"resident-graphql/middlewares"
	"resident-graphql/models/tables"
	"time"

	"github.com/graphql-go/graphql"
)

// CreateProvinceResolver func
func CreateProvinceResolver(p graphql.ResolveParams) (interface{}, error) {
	db := *connection.GetConnection()

	rand.Seed(time.Now().UnixNano())

	verifToken, err := middlewares.VerifyToken(helper.Token)
	if err != nil {
		return nil, err
	}
	if verifToken["role"] != "admin" && verifToken["role"] != "entry" {
		req := fmt.Sprintf("%v", verifToken["email"]) + " is trying to access " + helper.ReqBody
		middlewares.Sentry(req)
		return nil, err
	}

	var province tables.Provinces

	province.ID = uint(rand.Intn(100000))
	province.ProvinceName = p.Args["province_name"].(string)

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
	name, _ := p.Args["province_name"].(string)

	verifToken, err := middlewares.VerifyToken(helper.Token)
	if err != nil {
		return nil, err
	}
	if verifToken["role"] != "admin" && verifToken["role"] != "entry" {
		req := fmt.Sprintf("%v", verifToken["email"]) + " is trying to access " + helper.ReqBody
		middlewares.Sentry(req)
		return nil, err
	}

	var province tables.Provinces

	db.Model(&province).Where("id = ?", id).Update("province_name", name)

	return province, nil
}

// DeleteProvinceResolver func
func DeleteProvinceResolver(p graphql.ResolveParams) (interface{}, error) {
	db := *connection.GetConnection()
	id, _ := p.Args["id"].(int)

	verifToken, err := middlewares.VerifyToken(helper.Token)
	if err != nil {
		return nil, err
	}
	if verifToken["role"] != "admin" {
		req := fmt.Sprintf("%v", verifToken["email"]) + " is trying to access " + helper.ReqBody
		middlewares.Sentry(req)
		return nil, err
	}

	var province tables.Provinces

	db.Delete(&province, id)
	fmt.Println(province)

	return province, nil
}
