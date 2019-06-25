package controllers

import (
	"fmt"
	"kunpeng/budget-api/database"
	"kunpeng/budget-api/models"

	"github.com/kataras/iris"
	validator "gopkg.in/go-playground/validator.v9"
)

// GetFinance 获取全部的配置信息，按ctime排序
func GetFinance(ctx iris.Context) {
	// id := ctx.Params().GetString("id")
	financies, err := models.GetFinance()

	if err != nil && err.Error() != database.GetErrNotFound().Error() {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(APIResource(false, nil, err.Error()))
	} else {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(APIResource(true, financies, "success"))
	}
}

// CreateFinance 表示新建一条finance记录
func CreateFinance(ctx iris.Context) {
	finance := new(models.Finance)
	if err := ctx.ReadJSON(&finance); err != nil {
		ctx.StatusCode(iris.StatusMethodNotAllowed)
		ctx.JSON(errorData(err))
	} else {
		err := validate.Struct(finance)
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
			err := models.CreateFinance(finance)
			ctx.StatusCode(iris.StatusOK)
			if err != nil {
				ctx.JSON(APIResource(false, nil, "failed"))
			} else {
				ctx.JSON(APIResource(true, finance, "success"))
			}
		}
	}
}

// UpdateFinance 更新一条finance
func UpdateFinance(ctx iris.Context) {
	finance := new(models.Finance)
	if err := ctx.ReadJSON(&finance); err != nil {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(errorData(err))
	} else {
		err := validate.Struct(finance)
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
			err := models.UpdateFinance(id, finance)
			ctx.StatusCode(iris.StatusOK)
			if err != nil {
				ctx.JSON(APIResource(false, nil, "failed"))
			} else {
				ctx.JSON(APIResource(true, finance, "success"))
			}
		}
	}
}
