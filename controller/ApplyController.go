package controller

import (
	"Registration-system/common"
	"Registration-system/model"
	"Registration-system/response"
	"net/http"
	"Registration-system/service"
	"github.com/gin-gonic/gin"
)

func Apply(ctx *gin.Context) {
	DB:=common.GetDb()
	//获取参数
	applyInfo:=model.ApplyInfo{}
	ctx.ShouldBind(&applyInfo)
	// department1:=applyInfo.Department1
	// department2:=applyInfo.Department2
	// phoneNum:=applyInfo.PhoneNum
	// stuName:=applyInfo.StuName
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
	DB.Create(&applyInfo)
	response.Success(ctx,nil,"OK")
}