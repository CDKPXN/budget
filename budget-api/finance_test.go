package main

import (
	"fmt"
	"testing"

	// "gopkg.in/mgo.v2/bson"

	"github.com/kataras/iris"
)

var finance = map[string]interface{}{
	"_id":                      "5d1096aa37bcd528e88fe949",
	"management_cost":          1.6,
	"hydropower_property":      1.7,
	"travel":                   1.8,
	"entertain":                1.9,
	"salary":                   2.0,
	"social_security":          2.1,
	"public_accumulation_fund": 2.2,
}

func TestCreateFinance(t *testing.T) {
	expect := map[string]interface{}{
		"_id":                      finance["_id"],
		"management_cost":          finance["management_cost"],
		"hydropower_property":      finance["hydropower_property"],
		"travel":                   finance["travel"],
		"entertain":                finance["entertain"],
		"salary":                   finance["salary"],
		"social_security":          finance["social_security"],
		"public_accumulation_fund": finance["public_accumulation_fund"],
	}

	url := "/api/finance"
	create(t, url, finance, iris.StatusOK, true, "success", expect)
}

func TestGetFinance(t *testing.T) {
	// expect := map[string]interface{}{
	// 	"_id":                      finance["_id"],
	// 	"management_cost":          finance["management_cost"],
	// 	"hydropower_property":      finance["hydropower_property"],
	// 	"travel":                   finance["travel"],
	// 	"entertain":                finance["entertain"],
	// 	"salary":                   finance["salary"],
	// 	"social_security":          finance["social_security"],
	// 	"public_accumulation_fund": finance["public_accumulation_fund"],
	// }

	url := "/api/finance"
	getMore(t, url, iris.StatusOK, true, "success", nil)
}
func TestUpdateFinance(t *testing.T) {

	finance["management_cost"] = 2.1
	finance["hydropower_property"] = 2.2
	finance["travel"] = 2.3
	finance["entertain"] = 2.4
	finance["salary"] = 2.5
	finance["social_security"] = 2.6
	finance["public_accumulation_fund"] = 2.7

	expect := map[string]interface{}{
		"_id":                      finance["_id"],
		"management_cost":          finance["management_cost"],
		"hydropower_property":      finance["hydropower_property"],
		"travel":                   finance["travel"],
		"entertain":                finance["entertain"],
		"salary":                   finance["salary"],
		"social_security":          finance["social_security"],
		"public_accumulation_fund": finance["public_accumulation_fund"],
	}

	url := "/api/finance/%s"

	// fmt.Println(fmt.Sprintf(url, finance["_id"]))

	update(t, fmt.Sprintf(url, finance["_id"]), finance, iris.StatusOK, true, "success", expect)
}
