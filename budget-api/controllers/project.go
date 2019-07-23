package controllers

import (
	"fmt"
	"budget-api/database"
	"budget-api/models"

	"github.com/kataras/iris"
	validator "gopkg.in/go-playground/validator.v9"
)

// 5d0c8cd037bcd54078bfde55

// GetProject 获取全部项目信息
func GetProject(ctx iris.Context) {
	projects, err := models.GetProjects()

	if err != nil && err.Error() != database.GetErrNotFound().Error() {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(APIResource(false, nil, err.Error()))
	} else {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(APIResource(true, projects, "success"))
	}

}

// // GetProjectByID 获取一个项目的全部信息
// func GetProjectByID(ctx iris.Context) {

// }

// CreateProject 创建一个项目，创建时必须带上名称、周期、预计人员、预计人员工时
func CreateProject(ctx iris.Context) {
	project := new(models.Project)
	if err := ctx.ReadJSON(&project); err != nil {
		ctx.StatusCode(iris.StatusMethodNotAllowed)
		ctx.JSON(errorData(err))
	} else {
		err := validate.Struct(project)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println()
				fmt.Println(err.Namespace())
				fmt.Println(err.Field())
				fmt.Println(err.Type())
				fmt.Println(err.Param())
				fmt.Println()
			}
		} else {
			err := models.CreateProject(project)
			ctx.StatusCode(iris.StatusOK)
			if err != nil {
				ctx.JSON(APIResource(false, nil, "failed"))
			} else {
				ctx.JSON(APIResource(true, project, "success"))
			}
		}
	}
}

// DeleteProject 删除一个项目
func DeleteProject(ctx iris.Context) {
	id := ctx.Params().GetString("id")
	err := models.DeleteProject(id)
	ctx.StatusCode(iris.StatusOK)
	if err != nil {
		ctx.JSON(APIResource(false, nil, "failed"))
	} else {
		ctx.JSON(APIResource(true, nil, "success"))
	}
}

// UpdateProject 更新一个项目的基本信息
func UpdateProject(ctx iris.Context) {
	project := new(models.Project)
	if err := ctx.ReadJSON(&project); err != nil {
		ctx.StatusCode(iris.StatusMethodNotAllowed)
		ctx.JSON(errorData(err))
	} else {
		err := validate.Struct(project)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println()
				fmt.Println(err.Namespace())
				fmt.Println(err.Field())
				fmt.Println(err.Type())
				fmt.Println(err.Param())
				fmt.Println()
			}
		} else {
			id := ctx.Params().GetString("id")
			err := models.UpdateProject(id, project)
			ctx.StatusCode(iris.StatusOK)
			if err != nil {
				ctx.JSON(APIResource(false, nil, "failed"))
			} else {
				ctx.JSON(APIResource(true, project, "success"))
			}
		}
	}
}

// CreateProjectWorkload 创建时间——人员——工时组
func CreateProjectWorkload(ctx iris.Context) {

	tf := new(models.TimeStaffWorkload)
	if err := ctx.ReadJSON(&tf); err != nil {
		ctx.StatusCode(iris.StatusMethodNotAllowed)
		ctx.JSON(errorData(err))
	} else {
		err := validate.Struct(tf)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println()
				fmt.Println(err.Namespace())
				fmt.Println(err.Field())
				fmt.Println(err.Type())
				fmt.Println(err.Param())
				fmt.Println()
			}
		} else {
			id := ctx.Params().GetString("id")
			
			err := models.CreateProjectTimeStaffWorkload(id, tf)
			ctx.StatusCode(iris.StatusOK)
			if err != nil {
				ctx.JSON(APIResource(false, nil, "failed"))
			} else {
				ctx.JSON(APIResource(true, tf, "success"))
			}
		}
	}

}

// // UpdateProjectWorkload 修改时间——人员——工时组
// func UpdateProjectWorkload(ctx iris.Context) {

// }
