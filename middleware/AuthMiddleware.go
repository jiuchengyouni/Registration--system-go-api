package middleware

import (
	"Registration-system/common"
	entiy "Registration-system/entiy/parm"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString:=ctx.GetHeader("Authorization")
		if tokenString==""||!strings.HasPrefix(tokenString,"Bearer "){
			ctx.JSON(http.StatusUnauthorized,gin.H{
				"code":401,
				"msg":"权限不足",
			})
			ctx.Abort()//阻止调用后续的处理函数
			return
		}
		//去除前缀
		tokenString=tokenString[7:]
		token,claims,err:=common.ParseToken(tokenString)
		if err!=nil||!token.Valid{
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg": "权限不足",
			})
			ctx.Abort()
			return
		}
		//验证通过后获取claims中的userId
		userId:=claims.UserId
		DB:=common.GetDb()
		var user entiy.LoginInfo
		DB.First(&user,userId)
		//用户是否存在
		if user.StuNum==""{
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg": "权限不足",
			})
			ctx.Abort()
			return
		}
		//用户存在，将user信息写入上下文
		ctx.Set("user",user)
		ctx.Next()//调用后续处理函数
		//执行内部的pending handler
	}
}