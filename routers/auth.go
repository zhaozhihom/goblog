package router

import (
	"blog/config"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	var user config.User
	c.Bind(&user)
	logger.Infof("%+v", user)
	logger.Infof("%+v", users)
	for _, v := range users {
		if v == user {
			session := sessions.Default(c)
			session.Set("username", user.Username)
			session.Save()
			c.JSON(200, gin.H{"message": "登陆成功"})
			return
		}
	}
	c.JSON(401, gin.H{"message": "用户名密码有误"})
}

func checkIsLogin(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username")

	logger.Info("username:", username)

	if username == nil {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "需要登录",
		})
		return
	}

	c.AbortWithStatusJSON(200, gin.H{
		"message": "已登陆",
	})
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	// session.Set("dummy", "content")
	session.Delete("username")
	// session.Clear()

	// session.Options(sessions.Options{MaxAge: -1})
	session.Save()
	c.Redirect(301, "/")
}
