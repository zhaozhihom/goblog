package router

import (
	"blog/config"

	"github.com/gin-gonic/gin"
)

var logger = config.GetLogger()

// InitRouter 初始化路由
func InitRouter(conf *config.Config) {
	r := gin.Default()
	v1 := r.Group(conf.Server.BasePath)
	v1.POST("/post", postPost)
	v1.GET("/post/:id", getPost)
	v1.GET("/posts", getPosts)
	v1.PUT("/post", putPost)
	v1.DELETE("/post", deletePost)
	v1.PUT("/post/click", clickPost)
	v1.GET("/allpost", getAllPosts)
	v1.PUT("/post/markdeleted/:id", markDeletePost)

	r.Run(":" + conf.Server.Port)
}
