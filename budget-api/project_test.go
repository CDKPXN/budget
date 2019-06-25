package main

import (
	"fmt"
	"testing"

	"github.com/kataras/iris"
)

var planningWorkloads = []map[string]interface{}{
	{"staff_name": "zhangsan",
		"workload": 10},
	{"staff_name": "lisi",
		"workload": 20},
	{"staff_name": "wangwu",
		"workload": 15},
}

var realWorkloads = map[string]interface{}{
	// "time": time.Now(),
	"staff_workloads": []map[string]interface{}{
		{"staff_name": "zhangsan",
			"workload": 2},
		{"staff_name": "lisi",
			"workload": 3},
		{"staff_name": "wangwu",
			"workload": 5},
	},
}

var project = map[string]interface{}{
	"_id":                "5d1095cb37bcd5575cc53bc6",
	"planning_days":      120,
	"name":               "测试项目",
	"planning_workloads": planningWorkloads,
	"status":             1,
}

func TestCreateProject(t *testing.T) {
	expect := map[string]interface{}{
		"_id":                project["_id"],
		"name":               project["name"],
		"planning_days":      project["planning_days"],
		"planning_workloads": project["planning_workloads"],
		"status":             project["status"],
	}

	url := "/api/project"
	create(t, url, project, iris.StatusOK, true, "success", expect)
}

func TestGetProject(t *testing.T) {
	url := "/api/project"
	getMore(t, url, iris.StatusOK, true, "success", nil)
}
func TestUpdateProject(t *testing.T) {

	project["name"] = "changed"
	project["status"] = 2

	expect := map[string]interface{}{
		"_id":    project["_id"],
		"status": project["status"],
		"name":   project["name"],
	}

	url := "/api/project/%s"
	update(t, fmt.Sprintf(url, project["_id"]), project, iris.StatusOK, true, "success", expect)
}

func TestCreateProjectWorkload(t *testing.T) {
	// /{id:string}/workload
	expect := map[string]interface{}{
		// "time":            realWorkloads["time"],
		"staff_workloads": realWorkloads["staff_workloads"],
	}
	url := "/api/project/%s/workload"
	create(t, fmt.Sprintf(url, project["_id"]), realWorkloads, iris.StatusOK, true, "success", expect)
}

func TestDeleteProject(t *testing.T) {
	url := "/api/project/%s"
	delete(t, fmt.Sprintf(url, project["_id"]), iris.StatusOK, true, "success", nil)
}
