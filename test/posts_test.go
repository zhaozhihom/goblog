package test

import (
	"os"
	"testing"

	"blog/config"
	"blog/dao"

	"gopkg.in/yaml.v3"
)


func init() {
	log := config.GetLogger()

	var conf config.Config
	f, err := os.Open("../config.yaml")
	if err != nil {
		log.Error("读取配置文件错误, ", err)
	}
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&conf)
	if err != nil {
		log.Error("解析配置文件错误, ", err)
	}
	log.Debugf("读取到配置: %+v", conf)

	// 初始化数据库连接
	dao.InitDB(&conf)
}

// TestInsertPostFunc 测试新增文章
func TestInsertPost(t *testing.T) {
	count, err := dao.InsertPost(&dao.Posts{
		Title:   "测试",
		Content: "测试文章",
	})
	if count == 1 && err == nil {
		t.Log("测试通过！")
	} else {
		t.Error("测试失败! ", err)
	}
}
