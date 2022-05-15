package main

import (
	"context"
	"errors"
	"ginStudy/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	engine := service.SetEngine()

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	//_ = engine.Run()

	// 优雅关闭（go开启监听，监视关闭时候的信号，再进行shotdown）
	StartWithGracefulShowDown(engine)
}




func StartWithGracefulShowDown(engine *gin.Engine) {
	// 或者使用http启动
	srv := &http.Server{
		Addr:    ":8080",
		Handler: engine, // 实际上engine就是一个handler，从server传递到engine
	}

	// 在goroutine中开启server，这不会锁住graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n\n", err)
		}
	}()

	// 等待中断信号，五秒后优雅的shutdown服务器
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 等待信号传入quit
	<-quit                                               // 读出信号，否则一直堵塞
	log.Println("Shutting down server...")

	// 这个context用来引入
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 优雅地关闭
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting!")
}

