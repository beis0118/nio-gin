package Handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type BindJsonTest struct {
	Name string `json:"name"`
	Age string `json:"age"`
}

func BindJsonHandler(c *gin.Context) {
	var bindInfo BindJsonTest
	err := c.ShouldBindJSON(&bindInfo)// bindInfo自动被填写信息
	if err != nil {
		c.String(http.StatusInternalServerError,"你必须填写信息！")
	}
	c.String(http.StatusOK,"名字：%s, 年龄：%s",bindInfo.Name,bindInfo.Age)
}

func BindMustJsonHandler(c *gin.Context) {
	var bindInfo BindJsonTest
	err := c.MustBindWith(&bindInfo,binding.JSON)
	if err != nil {
		c.String(http.StatusInternalServerError,"你必须填写信息！")
	}
	c.String(http.StatusOK,"名字：%s, 年龄：%s",bindInfo.Name,bindInfo.Age)
}