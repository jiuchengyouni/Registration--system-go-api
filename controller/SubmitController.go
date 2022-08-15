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
// @Tags 提交信息接口
// @Summary 纳新报名
// @Description 填写报名信息，根据学号判断用户是否曾经报名过。
// @Accept json
// @Produce application/json
// @Param apply body entiyParm.ApplyInfo true "applyInfo"
// @Success 200 {string} string "{OK}"
// @Failure 404 {string} string "{Not Found}"
// @Router /student/apply [Post]
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
// @Tags 提交信息接口
// @Summary 登录
// @Description 确定用户学号。
// @Accept json
// @Produce application/json
// @Param apply body entiyParm.LoginInfo true "loginInfo"
// @Success 200 {string} string "{OK}"
// @Failure 404 {string} string "{Not Found}"
// @Router /student/login [Post]
func Login(ctx *gin.Context){
	DB:=common.GetDb()
	loginInfo:=entiyParm.LoginInfo{}
	ctx.BindJSON(&loginInfo)
	//数据验证
	//.....
	if !service.IsStuNum_login(DB,loginInfo.StuNum){
		response.Response(ctx, http.StatusUnprocessableEntity, 404, nil, "Not Found")
		return
	}else{
		response.Success(ctx,nil)
	}

	//jwt
	// token,err:=common.ReleaseToken(loginInfo)
	// if err!=nil{
	// 	response.Fail(ctx,nil,"系统异常")
	// 	log.Printf("token generate error : %v", err)
	// 	return
	// }
	//返回结果
	// response.Success(ctx,gin.H{"token":token})
}

// @Tags 提交信息接口
// @Summary 等待人数
// @Description 查看前面等待人数。
// @Accept json
// @Produce application/json
// @Param apply body entiyParm.WaitInfo true "waitInfo"
// @Success 200 {string} string "{OK}"
// @Failure 404 {string} string "{Not Found}"
// @Router /student/wait [Post]
func Wait(ctx *gin.Context){
	DB:=common.GetDb() 
	//获取参数
	var waitInfo entiyParm.WaitInfo
	ctx.ShouldBind(&waitInfo)
	number:=service.GetWaiter(DB,waitInfo.Department,waitInfo.StuNum)
	response.Success(ctx,gin.H{"wiaterNumber":number})
}

// @Tags 提交信息接口
// @Summary 面试预约
// @Description 根据学号判断用户是否已报名，已报名的才可以预约。。
// @Accept json
// @Produce application/json
// @Param apply body entiyParm.OrderInfo true "orderInfo"
// @Success 200 {string} string "{OK}"
// @Failure 404 {string} string "{Not Found}"
// @Router /admin/order [Post]
func Order(ctx *gin.Context){
	DB:=common.GetDb()
	//获取参数
	orderInfo:=entiyParm.OrderInfo{}
	ctx.ShouldBind(&orderInfo)
	//数据验证
	//.....
	//applyInfo:=service.IsOrder(DB,orderInfo.StuNum)
	// if applyInfo.Department1!=orderInfo.OrderDepartment&&applyInfo.Department2!=orderInfo.OrderDepartment{
	// 	response.Response(ctx, http.StatusUnprocessableEntity, 404, nil, "Not Found")
	// 	return
	// }
	//创建预约
	if service.Order(DB,orderInfo){
		response.Success(ctx,nil)
	}else {
		response.Response(ctx, http.StatusUnprocessableEntity, 404, nil, "Not Found")
	}
}

