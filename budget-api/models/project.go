package models

import (
	"budget-api/database"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// StaffWorkload 员工：工作量对
type StaffWorkload struct {
	StaffName string `json:"staff_name" bson:"staff_name" validate:"required,min=2,max=10"`
	Workload  uint   `json:"workload" bson:"workload" validate:"required,min=0"`
}

// TimeStaffWorkload 时间：员工：工作量
type TimeStaffWorkload struct {
	// ID             bson.ObjectId   `josn:"_id" bson:"_id"`
	Time           time.Time       `json:"time" bson:"time"`
	StaffWorkloads []StaffWorkload `json:"staff_workloads" bson:"staff_workloads" validate:"required"`
}

// Project 项目数据
type Project struct {
	ID                bson.ObjectId       `json:"_id" bson:"_id"`
	Ctime             time.Time           `json:"ctime" bson:"ctime"`
	Mtime             time.Time           `json:"mtime" bson:"mtime"`
	Name              string              `json:"name" bson:"name" validate:"required"`                             //项目名称
	Cost              uint64              `json:"cost" json:"cost" bson:"cost"`                                     //实际成本
	Budget            uint64              `json:"budget" bson:"budget"`                                             //预算
	StartTime         time.Time           `json:"start_time" bson:"start_time"`                                     //起始时间
	EndTime           time.Time           `json:"end_time" bson:"end_time"`                                         //结束时间
	PlanningDays      uint                `json:"planning_days" bson:"planning_days" validate:"required,min=0"`     //计划时间周期(天)
	PlanningWorkloads []StaffWorkload     `json:"planning_workloads" bson:"planning_workloads" validate:"required"` //计划员工工时
	RealWorkloads     []TimeStaffWorkload `json:"real_workloads" bson:"real_workloads"`                             //实际员工工时
	Status            uint                `json:"status" bson:"status" validate:"min=1,max=3"`                      //状态：1未开始；2正在进行；3完成
}

// CreateProject 创建一条项目记录
func CreateProject(project *Project) (err error) {
	project.Ctime = time.Now()
	project.Mtime = time.Now()
	project.Status = 1

	c := database.GetDataBase().C("project")
	err = c.Insert(&project)
	return
}

// GetProjects 获取所有项目记录
func GetProjects() (projects []Project, err error) {
	c := database.GetDataBase().C("project")
	if err = c.Find(nil).All(&projects); err != nil {
		return projects, err
	}
	return projects, nil
}

// DeleteProject 删除project记录
func DeleteProject(id string) (err error) {
	c := database.GetDataBase().C("project")

	if err = c.Remove(bson.M{"_id": bson.ObjectIdHex(id)}); err != nil {
		return
	}
	return
}

// UpdateProject 更新项目
func UpdateProject(id string, project *Project) (err error) {
	c := database.GetDataBase().C("project")

	var p Project
	// 如果项目没有，直接退
	if err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&p); err != nil {
		return
	}

	// 如果是更新状态
	if project.Status <= 0 || project.Status > 3 {
		return
	}
	if project.Status-p.Status != 1 {
		return
	}
	switch project.Status {
	case 2:
		project.StartTime = time.Now()
		break
	case 3:
		project.EndTime = time.Now()
		break
	default:
		break
	}

	err = c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{
		"name":               project.Name,
		"cost":               project.Cost,
		"budget":             project.Budget,
		"start_time":         project.StartTime,
		"end_time":           project.EndTime,
		"planning_days":      project.PlanningDays,
		"planning_workloads": project.PlanningWorkloads,
		"status":             project.Status}})
	return
}

// // UpdateProjectTimeStaffWorkload 更新项目实际人员工时
// func UpdateProjectTimeStaffWorkload(pid string, id string, tf *TimeStaffWorkload) (err error) {
// 	c := database.GetDataBase().C("project")

// 	var project Project

// 	if err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&project); err != nil {
// 		return
// 	}

// 	t := TimeStaffWorkload{time.Now(), StaffWorkloads}
// 	err = c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$addToSet": bson.M{"real_workloads": t, "mtime": time.Now()}})

// 	return
// }

// CreateProjectTimeStaffWorkload 项目实际人员工时
func CreateProjectTimeStaffWorkload(id string, tf *TimeStaffWorkload) (err error) {

	c := database.GetDataBase().C("project")

	var project Project

	if err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&project); err != nil {
		return
	}

	if tf.Time.IsZero() {
		tf.Time = time.Now()
	}

	err = c.Update(bson.M{"_id": bson.ObjectIdHex(id)},
		bson.M{"$addToSet": bson.M{"real_workloads": tf}})

	err = c.Update(bson.M{"_id": bson.ObjectIdHex(id)},
		bson.M{"$set": bson.M{"mtime": time.Now()}})

	return
}
