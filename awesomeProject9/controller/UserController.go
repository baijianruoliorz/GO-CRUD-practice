package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"oceanlearn.teach/ginessential/common"
	"oceanlearn.teach/ginessential/dto"
	"oceanlearn.teach/ginessential/model"
	"oceanlearn.teach/ginessential/response"
	"oceanlearn.teach/ginessential/util"
)

func Register(ctx *gin.Context) {
	DB:=common.GetDB()
	//获取参数
	name:=ctx.PostForm("name")
	telephone:=ctx.PostForm("telephone")
	password:=ctx.PostForm("possword")
	//数据验证
	if len(telephone)!=11{
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"mes":"手机号必须为11位"})
		return
	}
	if len(password)>6{
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"密码不得少于6位")
		return
	}
	//如果名称没传，则给它一个随机字符串
	if len(name)==0{
		name=util.RandomStiring(10)
	}
	log.Println(name,telephone,password)
	//判断手机号是否存在
	if isTelephoneExist(DB,telephone){
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"用户已经存在")
		return
	}
	//创建用户
	hasedPassword, err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err!=nil{
		response.Response(ctx,http.StatusUnprocessableEntity,500,nil,"加密错误")
		return
	}
	newUser:=model.User{
		Name: name,
		Telephone: telephone,
		Password: string(hasedPassword),
	}
	DB.Create(&newUser)
	response.Success(ctx,nil,"注册成功")
}
func Login(ctx *gin.Context){
	DB:=common.GetDB()
	//获取参数
	telephone:=ctx.PostForm("telephone")
	password:=ctx.PostForm("password")
	//数据验证
	if len(telephone)!=11{
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"手机号必须11位")
		return

	}
	if len(password)>6{
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"mes":"密码不得少于6位"})
		return
	}
	//判断手机号是否存在
	var user model.User
	DB.Where("telephone=?",telephone).First(&user)
	if user.ID==0{
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"用户不存在"})
		return
	}
	//判断密码是否正确
	if err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password));err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"code":400,"msg":"密码错误"})
		return
	}

	//发放token给前端
	token,err:=common.ReleaseToken(user)
	if err!=nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{"code":500,"msg":"系统异常"})
		log.Printf("token generate error: %v",err)
	}
	//返回结果
	ctx.JSON(200,gin.H{
		"code":200,
		"data":gin.H{"token":token},
		"msg":"登陆成功",
	})
	response.Success(ctx,gin.H{"token":token},"登陆成功")

}
func isTelephoneExist(db *gorm.DB,telephone string) bool{
	var user model.User
	db.Where("telephone=?",telephone).First(&user)
	if user.ID!=0{
		return true
	}
	return false
}
func Info(ctx *gin.Context){
	user,_:=ctx.Get("user")
	ctx.JSON(200,gin.H{"code":200,"data":gin.H{"user":dto.ToUserDto(user.(model.User))}})
}