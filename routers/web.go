package router

import (
	//"net/http"
	"blog/config"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var logger = config.GetLogger()

var users []config.User

// InitRouter 初始化路由
func InitRouter(conf *config.Config) {

	users = conf.Users

	r := gin.Default()

	//r.StaticFS("/blog", http.Dir("./build"))
	// r.StaticFS("/image", http.Dir("./build/image"))
	// r.StaticFile("/favicon.ico", "./build/favicon.ico")
	r.Use(static.Serve("/", static.LocalFile("./build", true)))

	r.NoRoute(func(c *gin.Context) {
		c.File("./build/index.html")
	})

	v1 := r.Group(conf.Server.BasePath)
	store := cookie.NewStore([]byte("username"))
	v1.Use(sessions.Sessions("mysession", store))

	public := v1.Group("/public")
	{
		public.POST("/login", login)
		public.GET("/checkAccess", checkIsLogin)
		public.POST("/logout", logout)
		public.GET("/posts", getPosts)
		public.GET("/allpost", getAllPosts)
		public.PUT("/post/click", clickPost)
		public.POST("/resume", postHTML)
	}

	private := v1.Group("/private")

	private.Use(authRequired())
	{
		private.POST("/post", postPost)
		private.GET("/post/:id", getPost)
		private.PUT("/post", putPost)
		private.DELETE("/post", deletePost)
		private.PUT("/post/markdeleted/:id", markDeletePost)
	}

	r.RunTLS(":443", "./tls.crt", "./tls.key")
	//r.Run(":" + conf.Server.Port)
}

func authRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		username := session.Get("username")

		logger.Info("authRequired username:", username)

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
