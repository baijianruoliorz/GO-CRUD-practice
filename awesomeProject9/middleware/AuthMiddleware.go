package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oceanlearn.teach/ginessential/common"
	"oceanlearn.teach/ginessential/model"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		tokenString:=ctx.GetHeader("Authorization")
		//validate token fomate
		if tokenString==" "||!strings.HasPrefix(tokenString,"Bearer"){
			ctx.JSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"权限不足"})
			ctx.Abort()
			return
		}
		tokenString=tokenString[7:]
		token,claims,err:=common.ParseToken(tokenString)
		if err!=nil||!token.Valid{
			ctx.JSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"权限不足"})
			ctx.Abort()
			return
		}
		//验证通过，获取claim中的userid
		userId:=claims.UserId
		DB:=common.GetDB()
		var user model.User
		DB.First(&user,userId)
		//用户
		if user.ID==0{
			ctx.JSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"权限不足"})
			ctx.Abort()
			return
		}
		//用户存在，user信息写上上下文
		ctx.Set("user",user)
		ctx.Next()
	}
}
