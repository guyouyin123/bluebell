package api

import (
	"bluebell/common"
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/sonwflake"
	"errors"
)

//用户注册的日志
func Register(p *models.ParamsUser) (err error) {
	//1.判断用户是否存在
	exist, err := mysql.ExistUser(p.Username)
	if err != nil {
		//数据库查询错误
		return err
	}
	if exist == false {
		return errors.New("用户已存在！！！")
	}
	//2.生成UID
	userID := sonwflake.GenID()
	p.UserID = userID                     //引用类型
	MD5Password := common.MD5(p.Password) //密码加密
	user := &models.User{
		UserId:   userID,
		UserName: p.Username,
		Password: MD5Password,
		Email:    p.Email,
		Gender:   p.Gender,
	}
	//3.保存进数据库
	err = mysql.SaveRegister(user)
	return err
}

func Login(p *models.LoginUser) (bool bool, err error) {
	p.Password = common.MD5(p.Password)
	//1.判断用户是否存在，密码是否正确
	exist, err := mysql.ExistLoginUser(p)
	if err != nil {
		//数据库查询错误
		return false, err
	}
	if exist == false {
		return false, nil
	}
	//2.验证通过，返回true,没有错误信息

	token, err2 := jwt.GenToken(p.UserID)
	if err2 != nil {
		return false, err
	}
	p.Token = token //引用类型，外层其他地方通过p拿到token
	return true, nil

}
