package routers

import (
	"bluebell/common"
	"bluebell/controllers"
	"bluebell/dao/mysql"
	_ "bluebell/docs"
	"bluebell/logger"
	"bluebell/pkg/jwt"
	"bluebell/pkg/sonwflake"
	"github.com/gin-gonic/gin"
	"net/http"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		//mode == release
		gin.SetMode(gin.ReleaseMode) //gin设置发布模式
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true)) //添加记录日志的中间件
	v1 := r.Group("/api/v1")

	//测试接口：需要登陆用户才能访问的接口,验证jwt.token
	//http://127.0.0.1:8080/api/v1/get_id/?userid=94372175745650688 header中还要携带token
	v1.GET("/get_id", jwt.JWTAuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"userid":   Test_id(),
			"username": common.GetTokenUserid(c),
		})
	})

	v1.POST("/register", controllers.UserRegister)

	v1.POST("/login", controllers.UserLogin)

	{
		v1.GET("/community", controllers.Community)
	}
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	// 其他请求
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}

func Test_id() int64 {
	id := sonwflake.GenID()
	return id
}

func Send_test() (*beautiful, error) {
	sqlStr := "select spu_id,title,price from dududu_shops where id = ?"
	var u beautiful
	err10 := mysql.Db.QueryRow(sqlStr, 101).Scan(&u.spu_id, &u.title, &u.price)
	if err10 != nil {
		return &u, err10
	}

	return &u, nil
}

type beautiful struct {
	spu_id string
	title  string
	price  string
}
