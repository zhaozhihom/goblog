package dao

import (
	"blog/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const createTable = `
	create table if not exists posts (
		id bigint(20) NOT NULL AUTO_INCREMENT,
		title varchar(100) not null,
		content text not null,
		click_times bigint(20) NOT NULL default 0,
		deleted tinyint(1) not null default 0,
		post_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (id)
	) DEFAULT CHARSET=utf8mb4;
`

var db *sqlx.DB

var logger = config.GetLogger()

// InitDB 初始化数据库连接
func InitDB(conf *config.Config) {
	var err error
	db, err = sqlx.Connect("mysql", conf.Database.Url)
	if err != nil {
		logger.Error("数据库连接异常:", err)
	}
	db.SetMaxIdleConns(conf.Database.MaxIdle)
	db.SetMaxOpenConns(conf.Database.MaxOpen)
	db.MustExec(createTable)
}

// GetDB 获取db对象
// func GetDB() (*sqlx.DB){
// 	return db
// }
