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

func Login(ctx *gin.Context){
	DB:=common.GetDb()
	loginInfo:=model.LoginInfo{}
	ctx.BindJSON(&loginInfo)
	stuNum:=loginInfo.StuNum
	//数据验证
	//.....
	if !service.IsStuNum_login(DB,stuNum){
		response.Response(ctx, http.StatusUnprocessableEntity, 404, nil, "Not Found")
		return
	}
	token,err:=common.ReleaseToken(loginInfo)
	if err!=nil{
		response.Fail(ctx,nil,"系统异常")
		log.Printf("token generate error : %v", err)
		return
	}
	//返回结果
	response.Success(ctx,gin.H{"token":token},"OK")
}