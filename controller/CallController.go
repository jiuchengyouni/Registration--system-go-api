package controller

import (
	"Registration-system/common"
	entiyParm "Registration-system/entiy/parm"
	entiyPo "Registration-system/entiy/po"
	"Registration-system/response"
	"Registration-system/service"
	"github.com/gin-gonic/gin"
)
// @Tags 管理员叫号接口
// @Summary 提前
// @Description 将用户排序提前一位。
// @Accept json
// @Produce application/json
// @Param changeInfo body entiyParm.ChangeInfo true "changeInfo"
// @Param department path string true "department"
// @Success 200 {string} string "{OK}"
// @Failure 404 {string} string "{Not Found}"
// @Router /admin/advance/{department} [Post]
func Advance(ctx *gin.Context) {
	DB:=common.GetDb()
	changeInfo:=entiyParm.ChangeInfo{}
	ctx.ShouldBind(&changeInfo)
	var wait entiyPo.Wait
	DB.Where("department=? AND stu_num=?",changeInfo.Department,changeInfo.StuNum).First(&wait)
	if service.Advance(DB,wait){
		response.Success(ctx,nil)
	}else{
		response.NotFound(ctx,nil)
	}
}
// @Tags 管理员叫号接口
// @Summary 删除
// @Description 删除用户排队。
// @Accept json
// @Produce application/json
// @Param changeInfo body entiyParm.ChangeInfo true "changeInfo"
// @Success 200 {string} string "{OK}"
// @Failure 404 {string} string "{Not Found}"
// @Router /admin/delete/{department} [Post]
func Delete(ctx *gin.Context){
	DB:=common.GetDb()
	changeInfo:=entiyParm.ChangeInfo{}
	ctx.ShouldBind(&changeInfo)
	var wait entiyPo.Wait
	DB.Where("department=? AND stu_num=?",changeInfo.Department,changeInfo.StuNum).First(&wait)
	if wait.StuName==""{
		response.NotFound(ctx,nil)
		return
	}
	DB.Delete(&wait)
	response.Success(ctx,nil)
}
// @Tags 管理员叫号接口
// @Summary 延后
// @Description 将用户排序延后一位。
// @Accept json
// @Produce application/json
// @Param changeInfo body entiyParm.ChangeInfo true "changeInfo"
// @Param department path string true "department"
// @Success 200 {string} string "{OK}"
// @Failure 404 {string} string "{Not Found}"
// @Router /admin/delay/{department} [Post]
func Delay(ctx *gin.Context){
	DB:=common.GetDb()
	changeInfo:=entiyParm.ChangeInfo{}
	ctx.ShouldBind(&changeInfo)
	var wait entiyPo.Wait
	DB.Where("department=? AND stu_num=?",changeInfo.Department,changeInfo.StuNum).First(&wait)
	if service.Delay(DB,wait){
		response.Success(ctx,nil)
	}else{
		response.NotFound(ctx,nil)
	}
}
// @Tags 管理员叫号接口
// @Summary 获取排队列表
// @Description 根据部门，获取对应排队列表。
// @Accept json
// @Produce application/json
// @Param department path string true "department"
// @Success 200 {string} string "{OK}"
// @Failure 404 {string} string "{Not Found}"
// @Router /admin/line/{department} [Get]
func Line(ctx *gin.Context){
	department:=ctx.Param("department")
	DB:=common.GetDb()
	waitline:=service.Waitline(DB,department)
	response.Success(ctx,gin.H{
		"data":waitline,
	})
}

// @Tags 管理员叫号接口
// @Summary 重置
// @Description 将用户从已结束状态从新插入队列（防止管理端误差）
// @Accept json
// @Produce application/json
// @Param waitInfo body entiyParm.WaitInfo true "waitInfo"
// @Success 200 {string} string "{OK}"
// @Failure 404 {string} string "{Not Found}"
// @Router /admin/reset [Post]
func Reset(ctx *gin.Context){
	DB:=common.DB
	waitInfo:=entiyParm.WaitInfo{}
	ctx.ShouldBind(&waitInfo)
	if service.Reset(DB,waitInfo){
		response.Success(ctx,nil);
	}else{
		response.NotFound(ctx,nil)
	}
}

// @Tags 管理员叫号接口
// @Summary 下一位
// @Description 根据部门，更新状态
// @Accept json
// @Produce application/json
// @Param department path string true "department"
// @Success 200 {string} string "{OK}"
// @Failure 404 {string} string "{Not Found}"
// @Router /admin/next/{department} [Get]
func Next(ctx *gin.Context){
	department:=ctx.Param("department")
	DB:=common.GetDb()
	if service.Next(DB,department){
		response.Success(ctx,nil)
	}else{
		response.NotFound(ctx,nil)
	}
}

// @Tags 管理员叫号接口
// @Summary 获取报名信息
// @Description 获取报名信息
// @Accept json
// @Produce application/json
// @Success 200 {string} string "{OK}"
// @Failure 404 {string} string "{Not Found}"
// @Router /admin/applyLine [Get]
func ApplyLine(ctx *gin.Context){
	DB:=common.GetDb()
	applyList:=service.GetApply(DB)
	response.Success(ctx,gin.H{
		"data":applyList,
	})
}