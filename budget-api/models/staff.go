package models

import (
	"kunpeng/budget-api/database"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Staff 职员
type Staff struct {
	ID     bson.ObjectId `json:"_id" bson:"_id"`
	Ctime  time.Time     `json:"ctime" bson:"ctime"`
	Mtime  time.Time     `json:"mtime" bson:"mtime"`
	Name   string        `json:"name" bson:"name" validate:"required,min=2,max=10"`
	Job    string        `json:"job" bson:"job" validate:"required,min=2,max=10"`
	Status int           `json:"status" bson:"status" validate:"min=1,max=2"`
	// Phone  string        `bson:"phone"`
	// Mail   string        `bson:"mail"`
}

// GetStaffByID 通过id获取staff记录
func GetStaffByID(id string) (staff Staff, err error) {
	c := database.GetDataBase().C("staff")
	if err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&staff); err != nil {
		if err.Error() != database.GetErrNotFound().Error() {
			return staff, err
		}
	}
	return staff, nil
}

// GetStaff 获取所有员工
func GetStaff() (staffs []Staff, err error) {
	c := database.GetDataBase().C("staff")
	if err = c.Find(nil).All(&staffs); err != nil {
		return staffs, err
	}
	return staffs, nil
}

// CreateStaff 创建用户
func CreateStaff(staff *Staff) (err error) {
	staff.Ctime = time.Now()
	staff.Mtime = time.Now()
	c := database.GetDataBase().C("staff")
	err = c.Insert(&staff)
	return
}

// UpdateStaff 更新员工信息
func UpdateStaff(id string, staff *Staff) (err error) {
	// id.Hex() 将bson.ObjectId转string
	c := database.GetDataBase().C("staff")
	err = c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"status": staff.Status, "mtime": time.Now()}})
	return
}
