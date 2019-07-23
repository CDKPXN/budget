package main

import (
	"flag"
	"budget-api/database"
	"budget-api/models"
	"os"
	"testing"

	"budget-api/config"

	"github.com/iris-contrib/httpexpect"
	"github.com/kataras/iris"
	"github.com/kataras/iris/httptest"
)

var (
	app         *iris.Application // iris.Applications
	testFinance *models.Finance
)

//单元测试基境
func TestMain(m *testing.M) {

	// 初始化app
	app = newApp()

	flag.Parse()
	exitCode := m.Run()

	database.GetDataBase().DropDatabase()

	os.Exit(exitCode)

}

// // 单元测试 login 方法
// func login(t *testing.T, url string, Object interface{}, StatusCode int, Status bool, Msg string, Data map[string]interface{}) (e *httpexpect.Expect) {
// 	e = httptest.New(t, app, httptest.Configuration{Debug: config.Conf.Get("app.debug").(bool)})
// 	if Data != nil {
// 		e.POST(url).WithJSON(Object).Expect().Status(StatusCode).JSON().Object().Values().Contains(Status, Msg, Data)
// 	} else {
// 		e.POST(url).WithJSON(Object).Expect().Status(StatusCode).JSON().Object().Values().Contains(Status, Msg)
// 	}

// 	return
// }

// 单元测试 create 方法
func create(t *testing.T, url string, Object interface{}, StatusCode int, Status bool, Msg string, Data map[string]interface{}) (e *httpexpect.Expect) {
	e = httptest.New(t, app, httptest.Configuration{Debug: config.Conf.Get("app.debug").(bool)})
	// at := GetLoginToken()

	ob := e.POST(url).WithJSON(Object).Expect().Status(StatusCode).JSON().Object()
	// ob := e.POST(url).WithHeader("Authorization", "Bearer "+at.Token).WithJSON(Object).
	// Expect().Status(StatusCode).JSON().Object()

	ob.Value("status").Equal(Status)
	ob.Value("msg").Equal(Msg)

	// id = ob.Value("data").Object().Value("_id").Raw().(string)

	for k, v := range Data {
		ob.Value("data").Object().Value(k).Equal(v)
	}

	return
}

// 单元测试 update 方法
func update(t *testing.T, url string, Object interface{}, StatusCode int, Status bool, Msg string, Data map[string]interface{}) (e *httpexpect.Expect) {
	e = httptest.New(t, app, httptest.Configuration{Debug: config.Conf.Get("app.debug").(bool)})

	ob := e.PUT(url).WithJSON(Object).Expect().Status(StatusCode).JSON().Object()

	ob.Value("status").Equal(Status)
	ob.Value("msg").Equal(Msg)

	for k, v := range Data {
		ob.Value("data").Object().Value(k).Equal(v)
	}

	return
}

// 单元测试 getOne 方法
func getOne(t *testing.T, url string, StatusCode int, Status bool, Msg string, Data map[string]interface{}) (e *httpexpect.Expect) {
	e = httptest.New(t, app, httptest.Configuration{Debug: config.Conf.Get("app.debug").(bool)})
	if Data != nil {
		e.GET(url).Expect().Status(StatusCode).JSON().Object().Values().Contains(Status, Msg, Data)
	} else {
		e.GET(url).Expect().Status(StatusCode).JSON().Object().Values().Contains(Status, Msg)
	}

	return
}

// 单元测试 getMore 方法
func getMore(t *testing.T, url string, StatusCode int, Status bool, Msg string, Data map[string]interface{}) (e *httpexpect.Expect) {
	e = httptest.New(t, app, httptest.Configuration{Debug: config.Conf.Get("app.debug").(bool)})
	if Data != nil {
		e.GET(url).Expect().Status(StatusCode).JSON().Object().Values().Contains(Status, Msg, Data)
	} else {
		e.GET(url).Expect().Status(StatusCode).JSON().Object().Values().Contains(Status, Msg)
	}

	return
}

// 单元测试 delete 方法
func delete(t *testing.T, url string, StatusCode int, Status bool, Msg string, Data map[string]interface{}) (e *httpexpect.Expect) {
	e = httptest.New(t, app, httptest.Configuration{Debug: config.Conf.Get("app.debug").(bool)})
	e.DELETE(url).Expect().Status(StatusCode).JSON().Object().Values().Contains(Status, Msg)
	return
}

/**
*登陆用户
*@return   Token 返回登陆后的token
 */
// func GetLoginToken() models.Token {
// 	response, status, msg := models.CheckLogin(
// 		config.Conf.Get("test.LoginUserName").(string),
// 		config.Conf.Get("test.LoginPwd").(string),
// 	)

// 	// 打印错误信息
// 	if !status {
// 		fmt.Println(msg)
// 	}

// 	return response
// }
