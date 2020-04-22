package router

import (
	"blog/config"
	"blog/dao"
	"blog/util"
	"strconv"

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

	r.Run(":" + conf.Server.Port)
}

func postPost(c *gin.Context) {
	var post dao.Posts
	c.Bind(&post)
	logger.Info("收到post:", post)
	_, err := dao.InsertPost(&post)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "success",
		})
	}
}

func getPosts(c *gin.Context) {
	page := c.Query("page")
	size := c.Query("size")
	if page == "" || size == "" {
		c.JSON(400, gin.H{
			"message": "缺少分页参数",
		})
		return
	}
	posts, err := dao.SelectPosts(util.CovertPageToOffset(page, size))
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "success",
			"data":    *posts,
		})
	}
}

func getPost(c *gin.Context) {
	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "参数错误",
		})
		return
	}
	post, err := dao.SelectPost(int64(ID))
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "success",
			"data":    *post,
		})
	}
}

func putPost(c *gin.Context) {
	var post dao.Posts
	c.Bind(&post)
	logger.Info("收到post:", post)
	if post.ID == 0 {
		c.JSON(400, gin.H{
			"message": "id不能为空",
		})
	}
	_, err := dao.UpdatePost(&post)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "success",
		})
	}
}

func deletePost(c *gin.Context) {
	var post dao.Posts
	c.Bind(&post)
	logger.Info("收到post:", post)
	if post.ID == 0 {
		c.JSON(400, gin.H{
			"message": "id不能为空",
		})
	}
	_, err := dao.DletePost(&post)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "success",
		})
	}
}

func clickPost(c *gin.Context) {
	var post dao.Posts
	c.Bind(&post)
	logger.Info("收到post:", post)
	if post.ID == 0 {
		c.JSON(400, gin.H{
			"message": "id不能为空",
		})
	}
	err := dao.UpdateClickTime(&post)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "success",
		})
	}
}
