package controller

import (
	"Registration-system/common"
	"Registration-system/model"
	"Registration-system/response"
	"Registration-system/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Advance(ctx *gin.Context) {
	department:=ctx.Param("department")
	DB:=common.GetDb()
	changeInfo:=model.ChangeInfo{}
	ctx.ShouldBind(&changeInfo)
	stuNum:=changeInfo.StuNum
	var waiterInfo model.WaitInfo
	var applyInfo model.ApplyInfo
	DB.Where("stu_num=? ",stuNum).First(&applyInfo)
	stuName:=applyInfo.StuName
	DB.Where("department=? AND stu_name=?",department,stuName).First(&waiterInfo)
	ID:=waiterInfo.ID
	var waiterInfo2 model.WaitInfo
	DB.Where("department=? AND id<?",department,ID).First(&waiterInfo2)
	waiterInfo2.StuName,waiterInfo.StuName=waiterInfo.StuName,waiterInfo2.StuName
	log.Print(waiterInfo)
	DB.Save(&waiterInfo)
	DB.Save(&waiterInfo2)
	response.Success(ctx,nil,"OK")
}

func Delete(ctx *gin.Context){
	department:=ctx.Param("department")
	DB:=common.GetDb()
	changeInfo:=model.ChangeInfo{}
	ctx.ShouldBind(&changeInfo)
	stuNum:=changeInfo.StuNum
	var waiterInfo model.WaitInfo
	var applyInfo model.ApplyInfo
	DB.Where("stu_num=? ",stuNum).First(&applyInfo)
	stuName:=applyInfo.StuName
	DB.Where("department=? AND stu_name=?",department,stuName).First(&waiterInfo)
	if waiterInfo.StuName==""{
		response.Response(ctx,http.StatusUnprocessableEntity,404,nil,"NotFound")
		return
	}
	DB.Delete(&waiterInfo)
	response.Success(ctx,nil,"OK")
}

func Delay(ctx *gin.Context){
	department:=ctx.Param("department")
	DB:=common.GetDb()
	changeInfo:=model.ChangeInfo{}
	ctx.ShouldBind(&changeInfo)
	stuNum:=changeInfo.StuNum
	var waiterInfo model.WaitInfo
	var applyInfo model.ApplyInfo
	DB.Where("stu_num=? ",stuNum).First(&applyInfo)
	stuName:=applyInfo.StuName
	DB.Where("department=? AND stu_name=?",department,stuName).First(&waiterInfo)
	ID:=waiterInfo.ID
	var waiterInfo2 model.WaitInfo
	DB.Where("department=? AND id>?",department,ID).First(&waiterInfo2)
	waiterInfo2.StuName,waiterInfo.StuName=waiterInfo.StuName,waiterInfo2.StuName
	DB.Save(&waiterInfo)
	DB.Save(&waiterInfo2)
	response.Success(ctx,nil,"OK")
}

func Line(ctx *gin.Context){
	department:=ctx.Param("department")
	DB:=common.GetDb()
	waitline:=service.Waitline(DB,department)
	response.Success(ctx,gin.H{
		"data":waitline,
	},"OK")
}