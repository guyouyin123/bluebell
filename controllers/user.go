package controllers

import (
	"bluebell/api"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

//用户注册方法
func UserRegister(c *gin.Context) {
	p := models.ParamsUser{}
	//1.参数校验 validator库
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("参数错误:", zap.Error(err))
		//判断err是不是validator中的类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			//c.JSON(http.StatusOK, gin.H{
			//	"code": 403,
			//	"msg":  err.Error(),
			//})
			ResponseWithMsg(c, CodeInvalidParam, err.Error())
			return
		}
		//c.JSON(http.StatusOK, gin.H{
		//	"msg":  "出错",
		//	"data": removeTopStruct(errs.Translate(trans)), //翻译
		//})
		ResponseWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	//2.用户注册接口
	if err := api.Register(&p); err != nil {
		zap.L().Error("用户已存在：", zap.Error(err))
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": "用户已存在！！！",
		//})
		ResponseError(c, CodeUserExist)
		return
	}

	//3.返回响应
	//c.JSON(http.StatusOK, gin.H{
	//	"msg":  "success",
	//	"data": p,
	//})

	ResponseSuccess(c, "注册成功！！！")

}

func UserLogin(c *gin.Context) {
	p := models.LoginUser{}
	//1.参数校验
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("绑定错误:", zap.Error(err))
		//判断err是不是validator类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			//不是validator中的错误类型
			//c.JSON(http.StatusOK, gin.H{
			//	"code": http.StatusOK,
			//	"msg":  err.Error(),
			//})
			ResponseError(c, CodeInvalidParam)
			return
		}
		//c.JSON(http.StatusOK, gin.H{
		//	"code": 403,
		//	"msg":  removeTopStruct(errs.Translate(trans)), //翻译validator中的错误
		//})
		ResponseWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//2.用户登陆接口
	exist, errs := api.Login(&p)
	if errs != nil {
		return
	}
	if exist == false {
		//c.JSON(403, gin.H{
		//	"code": 403,
		//	"msg":  "error",
		//	"data": fmt.Sprintf("%s用户名或密码错误！！！", p.Username),
		//})
		ResponseError(c, CodeInvalidPassword)
		return
	}

	//3.返回响应

	//c.JSON(http.StatusOK, gin.H{
	//	"code": http.StatusOK,
	//	"msg":  "success",
	//	"data": "登陆成功！！！",
	//})

	ResponseSuccess(c, "登陆成功！！！")

}
