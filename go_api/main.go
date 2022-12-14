package main

import (
	"context"
	"flag"
	"fmt"
	"go_api/utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go_api/app/config"
	"go_api/app/router"
	"go_api/settings"
	"go_api/utils/tcpServe"
)

func main() {
	var configName string
	flag.StringVar(&configName, "o", "settings.dev.yml", "环境配置")
	flag.Parse()
	config.Setup("settings/" + configName)

	r := router.Setup(config.ApplicationConfig)

	// 启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.ApplicationConfig.Host, config.ApplicationConfig.Port),
		Handler: r,
	}
	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			print(err)
			return
		}
	}()

	// 初始化数据
	settings.DataInitialize()
	// 连接Tcp终端
	if err := tcpServe.TcpClient(); err != nil {
		fmt.Println(err)
		return
	}
	// 获取token
	if err := utils.GetHttpToken(); err != nil {
		fmt.Println("获取 Token err!")
		fmt.Println(err)
		return
	}
	defer tcpServe.TcpConn.Close()

	fmt.Println("Server run at:")
	fmt.Printf("-  Local:   http://localhost:%s/ \r\n", config.ApplicationConfig.Port)
	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server Shutdown")
	}
}
