package routes

import (
	"gin_demo/dao/mysql"
	"gin_demo/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true)) //添加记录日志的中间件

	a, err := Send_test()

	if err != nil {
		panic(err)
	}

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, a.title)
	})

	return r
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
