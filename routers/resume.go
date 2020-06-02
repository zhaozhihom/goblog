package router

import (
	"blog/util"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type resume struct {
	Text string `form:"text"`
}

func postHTML(c *gin.Context) {
	var re resume
	c.Bind(&re)
	logger.Info(re)
	path, err := util.CovertHTMLToPdf(&re.Text)
	if err != nil {
		logger.Error(err)
		c.JSON(400, gin.H{
			"message": "转换错误",
		})
		return
	}
	paths := strings.Split(path, "/")
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", paths[len(paths)-1]))
    c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(path)
}