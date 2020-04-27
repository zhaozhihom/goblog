package router

import (
	"blog/config"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var logger = config.GetLogger()

// InitRouter 初始化路由
func InitRouter(conf *config.Config) {
	r := gin.Default()
	v1 := r.Group(conf.Server.BasePath)
	store := cookie.NewStore([]byte("username"))
	v1.Use(sessions.Sessions("mysession", store))

	public := v1.Group("/public")
	{
		public.GET("/posts", getPosts)
	}

	private := v1.Group("/private")
	
	private.Use(authRequired())
	{
		private.POST("/post", postPost)
		private.GET("/post/:id", getPost)
		private.PUT("/post", putPost)
		private.DELETE("/post", deletePost)
		private.PUT("/post/click", clickPost)
		private.GET("/allpost", getAllPosts)
		private.PUT("/post/markdeleted/:id", markDeletePost)
	}

	r.Run(":" + conf.Server.Port)
}

func authRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		username := session.Get("username")

		if username == nil {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "需要登录",
			})
			return
		}

		if username != "admin" {
			c.AbortWithStatusJSON(403, gin.H{
				"message": "没有权限",
			})
			return
		}

		c.Next()
	}
}
