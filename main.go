package main

import (
	"bluebell/controllers"
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/routers"
	"bluebell/settings"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title 这里写标题
// @version 1.0
// @description 这里写描述信息
// @termsOfService http://swagger.io/terms/

// @contact.name 这里写联系人信息
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 这里写接口服务的host
// @BasePath 这里写base path
func main() {
	//1.加载配置
	if err := settings.Init(); err != nil {
		fmt.Println("配置文件初始化失败：", err)
		return
	}

	//2.初始化日志
	// 方式一：viper配置文件方式
	if err := logger.Init(viper.GetString("mode")); err != nil {
		fmt.Println("日志文件初始化失败：", err)
		return
	}
	//方式二：结构体方式
	//if err := logger.Init2(settings.Conf.LogConfig); err != nil {
	//	fmt.Println("日志文件初始化失败：", err)
	//	return
	//}
	defer zap.L().Sync() //缓存日志追加到日志文件中
	//3.初始化mysql连接
	if err := mysql.Init(); err != nil {
		fmt.Println("mysql文件初始化失败：", err)
		return
	}
	defer mysql.Close()

	//4.初始化redis连接
	if err := redis.Init(); err != nil {
		fmt.Println("redis文件初始化失败：", err)
		return
	}
	defer redis.Close()

	// 翻译器初始化
	if err := controllers.InitTrans("zh"); err != nil {
		fmt.Println(err)
		zap.L().Error("controllers.InitTrans with err:", zap.Error(err))
	}

	//5.注册路由
	r := routers.Setup(viper.GetString("mode"))
	//6.启动服务

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")

}
