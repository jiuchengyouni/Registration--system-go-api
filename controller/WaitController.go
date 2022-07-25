package controller

import (
	"Registration-system/common"
	"Registration-system/model"
	"Registration-system/response"
	"Registration-system/service"

	"github.com/gin-gonic/gin"
)

func Wait(ctx *gin.Context){
	DB:=common.GetDb() 
	//获取参数
	var waitInfo model.WaitInfo
	ctx.ShouldBind(&waitInfo)
	department:=waitInfo.Department
	stuName:=waitInfo.StuName
	number:=service.GetWaiter(DB,department,stuName)
	response.Success(ctx,gin.H{"wiaterNumber":number},"OK")
}