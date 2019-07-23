package models

import (
	"fmt"
	"budget-api/database"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Finance 全局财务类配置数据
type Finance struct {
	ID                     bson.ObjectId `json:"_id" bson:"_id"`
	Ctime                  time.Time     `json:"ctime" bson:"ctime"`
	ManagementCost         float64       `json:"management_cost" bson:"management_cost" validate:"required"`                   //管理费用
	HydropowerProperty     float64       `json:"hydropower_property" bson:"hydropower_property" validate:"required"`           //物业水电
	Travel                 float64       `json:"travel" bson:"travel" validate:"required"`                                     //差旅
	Entertain              float64       `json:"entertain" bson:"entertain" validate:"required"`                               //招待
	Salary                 float64       `json:"salary" bson:"salary" validate:"required"`                                     //工资
	SocialSecurity         float64       `json:"social_security" bson:"social_security" validate:"required"`                   //社保
	PublicAccumulationFund float64       `json:"public_accumulation_fund" bson:"public_accumulation_fund" validate:"required"` //公积金

}

// CreateFinance 创建一条财务配置记录
func CreateFinance(finance *Finance) (err error) {
	finance.Ctime = time.Now()
	c := database.GetDataBase().C("finance")
	err = c.Insert(&finance)
	return
}

// GetFinance 获取全部finance记录
func GetFinance() (financies []Finance, err error) {
	c := database.GetDataBase().C("finance")
	//通过bson.M(是一个map[string]interface{}类型)进行
	if err = c.Find(nil).Sort("ctime").All(&financies); err != nil {
		if err.Error() != database.GetErrNotFound().Error() {
			return
		}
	}
	return
}

// UpdateFinance 更新一条finance记录
func UpdateFinance(id string, finance *Finance) (err error) {

	fmt.Println(id)

	fmt.Println(bson.ObjectIdHex(id))

	c := database.GetDataBase().C("finance")
	err = c.Update(bson.M{"_id": bson.ObjectIdHex(id)},
		bson.M{"$set": bson.M{
			"management_cost":          finance.ManagementCost,
			"hydropower_property":      finance.HydropowerProperty,
			"travel":                   finance.Travel,
			"entertain":                finance.Entertain,
			"salary":                   finance.Salary,
			"social_security":          finance.SocialSecurity,
			"public_accumulation_fund": finance.PublicAccumulationFund,
			"mtime":                    time.Now()}})
	return
}
