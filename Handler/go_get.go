package Handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	value := c.Value("domain")// 获取上下文值
	log.Println(value)
	c.HTML(http.StatusOK,"index.tmpl",gin.H{"title":"Hello!"})
}
func PingHandler(c *gin.Context) {
	c.JSON(200,gin.H{"message":"pong"})
}

func ParaHandler(c *gin.Context) {
	name := c.DefaultQuery("name","nioliu")
	age := c.Query("age")
	c.String(http.StatusOK,"name:%s , age:%s",name,age)
}

func TestHandler(c *gin.Context) {
	fullPath := c.FullPath()
	c.String(200,"这里是相对路径:%s\n",fullPath)
}

func V1Handler(c *gin.Context) {
		c.String(200,"这里是V1版本\n")
}

func PanicHandler(c *gin.Context) {
	// panic with a string -- the custom middleware could save this to a database or report it to the user
	panic("foo")
}

func CookieHandler(c *gin.Context) {
	cookie, err := c.Cookie("gin_cookie")

	if err != nil {
		cookie = "Not Set"
		c.SetCookie("gin_cookie","test",3600,"/","localhost",false,true)
	}

	log.Printf("Cookie value: %s \n", cookie)
}
