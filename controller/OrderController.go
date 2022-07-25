package controller

import (
	"Registration-system/common"
	"Registration-system/model"
	"Registration-system/response"
	"Registration-system/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Order(ctx *gin.Context){
	DB:=common.GetDb()
	//获取参数
	orderInfo:=model.OrderInfo{}
	ctx.ShouldBind(&orderInfo)
	orderDepartment:=orderInfo.OrderDepartment
	// phoneNum:=orderInfo.PhoneNum
	stuName:=orderInfo.StuName
	stuNum:=orderInfo.StuNum
	//数据验证
	//.....
	applyInfo:=service.IsOrder(DB,stuNum)
	if applyInfo.Department1!=orderDepartment&&applyInfo.Department2!=orderDepartment {
		response.Response(ctx, http.StatusUnprocessableEntity, 404, nil, "Not Found")
		return
	}
	//创建预约
	waitInfo:=model.WaitInfo{
		Department: orderDepartment,
		StuName: stuName,
	}
	DB.Create(&waitInfo)
	response.Success(ctx,nil,"OK")
}