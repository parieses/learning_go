package Router

import (
	"gin/Controllers"
	"gin/Middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	// 注册一个路由和处理函数
	engine := gin.Default()
	v1 := engine.Group("/v1",Middleware.Middleware1)
	{
		v1.GET("/login/:page", Controllers.LoginEndpoint)
		//v1.POST("/submit", submitEndpoint)
		//v1.POST("/read", readEndpoint)
	}
	engine.Any("/", Controllers.WebRoot)
	engine.GET("/user/:name/:qq", func(c *gin.Context) {
		name := c.Param("name")
		qq := c.Param("qq")

		c.String(http.StatusOK, "Hello %s %s", qq,name)
	})
	engine.GET("/users/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	return  engine
}
