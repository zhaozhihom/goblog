package main

import (
	"os"

	"blog/config"
	"blog/dao"
	router "blog/routers"

	"gopkg.in/yaml.v3"
)

func main() {

	log := config.GetLogger()

	var conf config.Config
	f, err := os.Open("./config.yaml")
	if err != nil {
		log.Error("读取配置文件错误, ", err)
		return
	}
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&conf)
	if err != nil {
		log.Error("解析配置文件错误, ", err)
		return
	}
	log.Debugf("读取到配置: %+v", conf)

	// 初始化数据库连接
	dao.InitDB(&conf)

	// 初始化路由
	router.InitRouter(&conf)
}
