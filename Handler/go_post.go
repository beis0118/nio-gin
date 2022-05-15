package Handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func FormPost(c *gin.Context) {
	message := c.PostForm("message")
	age := c.DefaultPostForm("age", "22")

	c.JSON(http.StatusOK,gin.H{"age":age,"message":message})
}
