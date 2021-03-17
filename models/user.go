package models

type User struct {
	UserId   int64  `Db:"user_id"`
	UserName string `Db:"username"`
	Password string `Db:"password"`
	Email    string `Db:"email"`
	Gender   string `Db:"gender"`
}
