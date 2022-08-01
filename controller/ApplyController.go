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
func Apply(ctx *gin.Context) {
	DB:=common.GetDb()
	//获取参数
	applyInfo:=entiyParm.ApplyInfo{}
	ctx.ShouldBind(&applyInfo)
	department1:=applyInfo.Department1
	department2:=applyInfo.Department2
	phoneNum:=applyInfo.PhoneNum
	stuName:=applyInfo.StuName
	stuNum:=applyInfo.StuNum
	//数据验证
	//不知道需要不要写,还是前端处理好了
	//.......
	//判断学号是否存在
	if service.IsStuNum(DB,stuNum){
		response.Response(ctx,http.StatusUnprocessableEntity,201,nil,"Created")
		return
	}
	//学号要加密吗?
	
	//创建用户
	// newApplyInfo:=model.ApplyInfo{
	// 	Department1: department1,
	// 	Department2: department2,
	// 	PhoneNum: phoneNum,
	// 	StuName: stuName,
	// 	StuNum: stuNum,
	// }
	newApply:=entiyPo.Apply{
		Department1: department1,
		Department2: department2,
		PhoneNum: phoneNum,
		StuName: stuName,
		StuNum: stuNum,
	}
	DB.Create(&newApply)
	response.Success(ctx,nil)
}