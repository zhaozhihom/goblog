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
		PRIMARY KEY (id)
	)
`

// InitDB 初始化数据库连接
func InitDB(conf *config.Config) {

	log := config.GetLogger()

	db, err := sqlx.Connect("mysql", conf.Database.Url)
	if err != nil {
		log.Error("数据库连接异常:", err)
	}
	db.SetMaxIdleConns(conf.Database.MaxIdle)
	db.SetMaxOpenConns(conf.Database.MaxOpen)
	db.MustExec(createTable)
}
