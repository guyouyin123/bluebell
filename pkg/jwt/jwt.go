package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

const TokenExpireDuration = time.Hour * 2 //token过期时间 2小时
var MySecret = []byte("token加盐，加盐")       //token加盐
var UserID = "userid"

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录两个username，userid字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID int64 `json:"userid"`
	//Username string `json:"Username"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(userid int64) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		//userID,   // 自定义字段
		userid, // 自定义字段
		jwt.StandardClaims{ //标准字段
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "bluebell-jeff",                            // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	var mc = new(MyClaims) //初始化内存
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里的具体实现方式要依据你的实际业务情况决定
		var QueryUserid, errs = strconv.ParseInt(c.Query("userid"), 10, 64) //url中携带的userid,转为int64
		if errs != nil {
			//请求参数错误
			c.JSON(http.StatusOK, gin.H{
				"code": 2002,
				"msg":  "请求参数错误",
			})
			c.Abort()
			return
		}

		authHeader := c.Request.Header.Get("token")
		if authHeader == "" {
			//可以使用封装的common.ResponseError()来返回
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}

		mc, err := ParseToken(authHeader)
		if err != nil || mc.UserID != QueryUserid {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上，记录登陆的用户名
		c.Set(UserID, mc.UserID)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}

func test_token() {
	aa, err := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjE2MDY5MDU0LCJpc3MiOiJibHVlYmVsbC1qZWZmIn0.2dom1qLq-MpJAY_ctg2a2h_Tlzx_t7kZBKi8lQuxUPw")
	fmt.Println(aa)
	fmt.Println(err)
}
