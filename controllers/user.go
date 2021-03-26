package controllers

import (
	"bluebell/api"
	"bluebell/common"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

//用户注册登陆逻辑
//用户注册方法

// GetPostListHandler2 用户注册接口
// @Summary 用户注册接口
// @Tags 用户注册相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamsUser false "查询参数"
// @Security ApiKeyAuth
// @Router /posts2 [get]

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
			common.ResponseWithMsg(c, common.CodeInvalidParam, err.Error())
			return
		}
		//c.JSON(http.StatusOK, gin.H{
		//	"msg":  "出错",
		//	"data": removeTopStruct(errs.Translate(trans)), //翻译
		//})
		common.ResponseWithMsg(c, common.CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	//2.用户注册接口
	if err := api.Register(&p); err != nil {
		zap.L().Error("用户已存在：", zap.Error(err))
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": "用户已存在！！！",
		//})
		common.ResponseError(c, common.CodeUserExist)
		return
	}

	//3.返回响应
	//c.JSON(http.StatusOK, gin.H{
	//	"msg":  "success",
	//	"data": p,
	//})

	common.ResponseSuccess(c, p)

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
			common.ResponseError(c, common.CodeInvalidParam)
			return
		}
		//c.JSON(http.StatusOK, gin.H{
		//	"code": 403,
		//	"msg":  removeTopStruct(errs.Translate(trans)), //翻译validator中的错误
		//})
		common.ResponseWithMsg(c, common.CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
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
		common.ResponseError(c, common.CodeInvalidPassword)
		return
	}

	//3.返回响应

	//c.JSON(http.StatusOK, gin.H{
	//	"code": http.StatusOK,
	//	"msg":  "success",
	//	"data": "登陆成功！！！",
	//})

	common.ResponseSuccess(c, p.Token)

}
