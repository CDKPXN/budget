package controllers

import (
	"fmt"
	"kunpeng/budget-api/database"
	"kunpeng/budget-api/models"

	"github.com/kataras/iris"
	validator "gopkg.in/go-playground/validator.v9"
)

// GetStaff 获取全部职员信息
func GetStaff(ctx iris.Context) {
	staffs, err := models.GetStaff()

	if err != nil && err.Error() != database.GetErrNotFound().Error() {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(APIResource(false, nil, err.Error()))
	} else {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(APIResource(true, staffs, "success"))
	}
}

// CreateStaff 创建职员
func CreateStaff(ctx iris.Context) {
	staff := new(models.Staff)
	if err := ctx.ReadJSON(&staff); err != nil {
		ctx.StatusCode(iris.StatusMethodNotAllowed)
		ctx.JSON(errorData(err))
	} else {
		err := validate.Struct(staff)
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
			err := models.CreateStaff(staff)
			ctx.StatusCode(iris.StatusOK)
			if err != nil {
				ctx.JSON(APIResource(false, nil, "failed"))
			} else {
				ctx.JSON(APIResource(true, staff, "success"))
			}
		}
	}
}

// UpdateStaff 更新职员
func UpdateStaff(ctx iris.Context) {
	staff := new(models.Staff)
	if err := ctx.ReadJSON(&staff); err != nil {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(errorData(err))
	} else {
		err := validate.Struct(staff)
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
			err := models.UpdateStaff(id, staff)
			ctx.StatusCode(iris.StatusOK)
			if err != nil {
				ctx.JSON(APIResource(false, nil, "failed"))
			} else {
				ctx.JSON(APIResource(true, staff, "success"))
			}
		}
	}
}
