package models

//type ParamsUser struct {
//	Username   string `json:"username" binding:"required"`
//	Password   string `json:"password" binding:"required"`
//	Repassword string `json:"re_password" binding:"required,eqfield=Password"`
//}

type ParamsUser struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Repassword string `json:"re_password" binding:"required,eqfield=Password"`
	Email      string `json:"email" binding:"required,email"`
	Gender     string `json:"gender" binding:"required"`
	UserID     int64
}

type LoginUser struct {
	UserID   int64  `json:"userid" binding:"required"`
	Password string `json:"password" binding:"required"`
	Token    string
}
