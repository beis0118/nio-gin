package Handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func MyCustomRecovery(c *gin.Context, recovered interface{}) {
	if err,ok := recovered.(string);ok {
		c.String(http.StatusInternalServerError,fmt.Sprintf("error: %s", err))
	}
	c.AbortWithStatus(http.StatusInternalServerError)
}

func WriteLogToFile() {
	//gin.DisableConsoleColor()// 关闭console颜色

	file, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(file)

	// 如果既写入日志，又console打印，则：
	gin.DefaultWriter = io.MultiWriter(os.Stdout,file)// multi为复制
}


func CustomLogFormat(param  gin.LogFormatterParams) string {
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}

/**
实际上就是生命周期
 */

func CustomMiddle(c *gin.Context) {
	t := time.Now()

	// 设置自定义变量
	c.Set("domain","tencent")

	// 请求前
	log.Println("请求前！！！")
	c.Next() // 这一步就是执行每个handler里的方法
	// 请求后
	log.Println("请求后！！！")
	latency := time.Since(t)
	log.Println(latency)

	// 访问发送状态
	status := c.Writer.Status()
	log.Println(status)
}

func CustomMiddle2(c *gin.Context) {
	t := time.Now()

	// 设置自定义变量
	c.Set("domain","tencent2")

	// 请求前
	log.Println("请求前2！！！")
	c.Next() // 这一步就是执行每个handler里的方法
	// 请求后
	log.Println("请求后2！！！")
	latency := time.Since(t)
	log.Println(latency)

	// 访问发送状态
	status := c.Writer.Status()
	log.Println(status)
}