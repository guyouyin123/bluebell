package routers

import (
	"bluebell/controllers"
	"bluebell/dao/mysql"
	"bluebell/logger"
	"bluebell/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		//mode == release
		gin.SetMode(gin.ReleaseMode) //gin设置发布模式
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true)) //添加记录日志的中间件

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "okok")
	})
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"user_id": Test_id(),
		})
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	r.POST("/register", controllers.UserRegister)

	r.POST("/login", controllers.UserLogin)
	return r
}

func Test_id() int64 {
	id := pkg.GenID()
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
