package controller

import (
	"Registration-system/common"
	entiyParm "Registration-system/entiy/parm"
	entiyPo "Registration-system/entiy/po"
	"Registration-system/response"
	"Registration-system/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Order(ctx *gin.Context){
	DB:=common.GetDb()
	//获取参数
	orderInfo:=entiyParm.OrderInfo{}
	ctx.ShouldBind(&orderInfo)
	//数据验证
	//.....
	applyInfo:=service.IsOrder(DB,orderInfo.StuNum)
	if applyInfo.Department1!=orderInfo.OrderDepartment&&applyInfo.Department2!=orderInfo.OrderDepartment{
		response.Response(ctx, http.StatusUnprocessableEntity, 404, nil, "Not Found")
		return
	}
	//创建预约
	wait:=entiyPo.Wait{
		Department: orderInfo.OrderDepartment,
		StuName: orderInfo.StuName,
		StuNum: orderInfo.StuNum,
		PhoneNum: orderInfo.PhoneNum,
	}
	DB.Create(&wait)
	response.Success(ctx,nil)
}