package database

import (
	"fmt"
	"kunpeng/budget-api/config"

	"gopkg.in/mgo.v2"
)

var session *mgo.Session
var database *mgo.Database

// Init 用于新建数据库连接
func Init() {
	connect := config.Conf.Get("mongodb.connect").(string)
	db := config.Conf.Get("mongodb.database").(string)

	fmt.Println(connect)

	session, err := mgo.Dial(connect)
	if err != nil {
		panic(fmt.Sprintf("No error should happen when connecting to  database, but got err=%+v", err))
	}

	session.SetMode(mgo.Monotonic, true)
	//使用指定数据库
	database = session.DB(db)

}

// func Close() {
// 	session.Close()
// }

// GetMgo 返回session
func GetMgo() *mgo.Session {
	return session
}

// GetDataBase 返回database
func GetDataBase() *mgo.Database {
	return database
}

// GetErrNotFound 返回错误
func GetErrNotFound() error {
	return mgo.ErrNotFound
}
