package mysql

import (
	"bluebell/models"
)

//用户注册入库
func SaveRegister(user *models.User) (err error) {
	sqlStr := `insert into user(user_id,username,password,email,gender) values(?,?,?,?,?)`
	_, err = Db.Exec(sqlStr, user.UserId, user.UserName, user.Password, user.Email, user.Gender)
	return err
}

//判断用户是否存在
func ExistUser(username string) (bool, error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := Db.Get(&count, sqlStr, username); err != nil {
		//false,用户已存在
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return true, nil
}

//判断用户是否存在,和密码是否正确
func ExistLoginUser(p *models.LoginUser) (bool, error) {
	sqlStr := `select count(user_id) from user where user_id = ? and password = ?`
	var count int
	if err := Db.Get(&count, sqlStr, p.UserID, p.Password); err != nil {
		//false,用户已存在
		return false, err
	}
	if count > 0 {
		//true用户名和密码验证通过
		return true, nil
	}
	return false, nil
}
