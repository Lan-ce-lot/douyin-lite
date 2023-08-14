package router

import (
	"github.com/gin-gonic/gin"
	"github.com/houqingying/douyin-lite/handler/comment"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"code":    200,
		})
	})
	//r.Static("static", config.Global.StaticSourcePath)
	baseGroup := r.Group("/douyin")
	baseGroup.POST("/comment/action/", comment.Action)
	baseGroup.GET("/comment/list/", comment.List)
	return r
}
