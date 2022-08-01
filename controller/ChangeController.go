package controller

import (
	"Registration-system/common"
	entiyParm "Registration-system/entiy/parm"
	entiyPo "Registration-system/entiy/po"
	"Registration-system/response"
	"Registration-system/service"
	"github.com/gin-gonic/gin"
)

func Advance(ctx *gin.Context) {
	department:=ctx.Param("department")
	DB:=common.GetDb()
	changeInfo:=entiyParm.ChangeInfo{}
	ctx.ShouldBind(&changeInfo)
	var wait entiyPo.Wait
	DB.Where("department=? AND stu_num=?",department,changeInfo.StuNum).First(&wait)
	if service.Advance(DB,wait){
		response.Success(ctx,nil)
	}else{
		response.NotFound(ctx,nil)
	}

}

func Delete(ctx *gin.Context){
	department:=ctx.Param("department")
	DB:=common.GetDb()
	changeInfo:=entiyParm.ChangeInfo{}
	ctx.ShouldBind(&changeInfo)
	var wait entiyPo.Wait
	DB.Where("department=? AND stu_num=?",department,changeInfo.StuNum).First(&wait)
	if wait.StuName==""{
		response.NotFound(ctx,nil)
		return
	}
	DB.Delete(&wait)
	response.Success(ctx,nil)
}

func Delay(ctx *gin.Context){
	department:=ctx.Param("department")
	DB:=common.GetDb()
	changeInfo:=entiyParm.ChangeInfo{}
	ctx.ShouldBind(&changeInfo)
	var wait entiyPo.Wait
	DB.Where("department=? AND stu_num=?",department,changeInfo.StuNum).First(&wait)
	if service.Delay(DB,wait){
		response.Success(ctx,nil)
	}else{
		response.NotFound(ctx,nil)
	}
}

func Line(ctx *gin.Context){
	department:=ctx.Param("department")
	DB:=common.GetDb()
	waitline:=service.Waitline(DB,department)
	response.Success(ctx,gin.H{
		"data":waitline,
	})
}