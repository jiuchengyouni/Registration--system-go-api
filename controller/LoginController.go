package controller

import (
	"Registration-system/common"
	entiyParm "Registration-system/entiy/parm"
	"Registration-system/response"
	"Registration-system/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

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