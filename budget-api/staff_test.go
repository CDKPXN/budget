package main

import (
	"fmt"
	"testing"

	// "gopkg.in/mgo.v2/bson"

	"github.com/kataras/iris"
)

var staff = map[string]interface{}{
	"_id":    "5d10966b37bcd55434e92212",
	"name":   "weizijie",
	"job":    "developer",
	"status": 1,
}

func TestCreateStaff(t *testing.T) {
	expect := map[string]interface{}{
		"_id":    staff["_id"],
		"name":   staff["name"],
		"job":    staff["job"],
		"status": staff["status"],
	}

	url := "/api/staff"
	create(t, url, staff, iris.StatusOK, true, "success", expect)
}

func TestGetStaff(t *testing.T) {
	// expect := map[string]interface{}{
	// 	"_id":    staff["_id"],
	// 	"name":   staff["name"],
	// 	"job":    staff["job"],
	// 	"status": staff["status"],
	// }

	url := "/api/staff"
	getMore(t, url, iris.StatusOK, true, "success", nil)
}
func TestUpdateStaff(t *testing.T) {

	staff["job"] = "coder"

	expect := map[string]interface{}{
		"_id":  staff["_id"],
		"job":  staff["job"],
		"name": staff["name"],
	}

	url := "/api/staff/%s"
	update(t, fmt.Sprintf(url, staff["_id"]), staff, iris.StatusOK, true, "success", expect)
}
