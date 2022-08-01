package controller

import (
	"Registration-system/common"
	entiyParm "Registration-system/entiy/parm"
	"Registration-system/response"
	"Registration-system/service"

	"github.com/gin-gonic/gin"
)

func Wait(ctx *gin.Context){
	DB:=common.GetDb() 
	//获取参数
	var waitInfo entiyParm.WaitInfo
	ctx.ShouldBind(&waitInfo)
	number:=service.GetWaiter(DB,waitInfo.Department,waitInfo.StuName)
	response.Success(ctx,gin.H{"wiaterNumber":number})
}