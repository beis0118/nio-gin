package service

import (
	"ginStudy/Handler"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func SetEngine() *gin.Engine {
	// 创建engine
	//engine := gin.New()
	//engine.Use(gin.Logger()) // 中间件其实就是handler，但是有生命周期
	//engine.Use(gin.Recovery())

	// 等同于上面： Default With the Logger and Recovery middleware already attached
	//engine = gin.Default()

	// 自定义一些middleware，必须紧跟着，否则无效
	engine := gin.New()
	// 日志写入文件, 这是全局设置
	Handler.WriteLogToFile()
	// 日志指定格式
	engine.Use(gin.LoggerWithFormatter(Handler.CustomLogFormat))
	// 自定义recovery
	engine.Use(gin.CustomRecovery(Handler.MyCustomRecovery))
	engine.Use(Handler.CustomMiddle)
	engine.Use(Handler.CustomMiddle2) // 入则从前往后执行中间件；出则从后往前执行中间件
	//_ = engine.SetTrustedProxies(nil) // 设置信任的代理，默认为全部（不安全）

	// 路由
	engine.GET("/panic", Handler.PanicHandler) // curl localhost:8080/panic
	engine.GET("/ping", Handler.PingHandler)   // http://localhost:8080/ping
	engine.GET("/para", Handler.ParaHandler)   // curl localhost:8080/para\?age=22\&name=qwdq
	// handler可以添多个，按顺序执行（实际上就是合并到一起，因为只有全部执行完后才会真正返回）
	engine.POST("/form", Handler.FormPost, Handler.PingHandler) // curl localhost:8080/form -F "age=23" -F "message=nindiwni"

	// 路由group
	routerGroup := engine.Group("/test") // 创建route组，这里指定相对路径
	routerGroup.Use(Handler.TestHandler) // 对所有test请求都添加这个handler，都会先执行
	{
		routerGroup.GET("/para", Handler.ParaHandler) // curl localhost:8080/test/para\?age=22\&name=qwdq

		v1Group := routerGroup.Group("/v1", Handler.V1Handler) // 群组可以嵌套，在后面可以直接添加handler
		v1Group.GET("/", Handler.PingHandler)                  // curl localhost:8080/test/v1/
	}

	bindGroup := engine.Group("/bind")
	{
		bindGroup.POST("/json", Handler.BindJsonHandler)          // curl -X POST -H "Content-Type:application/json" -d '{"name":"nioliu","age":"12"}' localhost:8080/bind/json
		bindGroup.POST("/must/json", Handler.BindMustJsonHandler) // curl -X POST -H "Content-Type:application/json" localhost:8080/bind/must/json
	}

	// 返回html页面
	engine.LoadHTMLGlob("../templates/*") // 指定全局模版路径
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html") // 指定模版
	engine.GET("/", Handler.IndexHandler) // curl localhost:8080

	// 重定向
	engine.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com/") // curl localhost:8080/redirect -L
		engine.HandleContext(c)                                           // 传递context
	})

	// 在handle、middleware开启goroutines时，要copy一个context
	engine.GET("/long_async", func(c *gin.Context) {
		// create copy to be used inside the goroutine
		// c则在外goroutine结束后关闭，而cCp没有关闭
		cCp := c.Copy()
		go func() { // 这里面的程序就不受外部的context影响，所以不会受外部的middleware处理
			// simulate a long task with time.Sleep(). 5 seconds
			time.Sleep(5 * time.Second)

			// note that you are using the copied context "cCp", IMPORTANT
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})

	engine.GET("/long_sync", func(c *gin.Context) {
		// simulate a long task with time.Sleep(). 5 seconds
		time.Sleep(5 * time.Second)

		// since we are NOT using a goroutine, we do not have to copy the context
		log.Println("Done! in path " + c.Request.URL.Path)
	})

	// 设置cookie
	engine.GET("/cookie", Handler.CookieHandler)
	return engine
}
