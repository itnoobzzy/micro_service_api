package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	server := http.Server{}

	// 把服务器放到 goroutine 中去执行
	go func() {
		if err := server.ListenAndServe(); err != nil {
			// 手动调用 Shutdown 时会产生这个错误
			if err != http.ErrServerClosed {
				log.Println(err)
			}
		}
	}()

	log.Printf("server pid: %d\n", os.Getpid())

	// 捕捉指定的信号
	quit := make(chan os.Signal)
	// 前台时，按 ^C 时触发
	signal.Notify(quit, syscall.SIGINT)
	// 后台时，kill 时触发。kill -9 时的信号 SIGKILL 不能捕捉，所以不用添加
	signal.Notify(quit, syscall.SIGTERM)

	// 等待退出信号
	sig := <-quit
	log.Printf("received signal: %v\n", sig)

	// 收到信号后，优雅地关闭服务器
	log.Println("server shutting down")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Println(err)
	}
	log.Println("server shutted down")
}
