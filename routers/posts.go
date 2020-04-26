package router

import (
	"blog/dao"
	"blog/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func getAllPosts(c *gin.Context) {
	posts, err := dao.SelectAllPosts()
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

func markDeletePost(c *gin.Context) {
	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "参数错误",
		})
		return
	}
	_, err = dao.MarkDletePost(int64(ID))
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
