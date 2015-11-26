package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//数据库的链接
const DB_NAME = "ruoli:ruoli@/beego?charset=utf8"

//初始化数据库
func RegisterDB() {
	// register driver
	// mysql sqlite2 postgres 这三种是已经注册过的，可以不用设置
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	// set default database
	// 必须注册一个别名为 default 的数据库，作为默认使用
	orm.RegisterDataBase("default", "mysql", DB_NAME, 30)
	// register model
	// create table
	orm.RunSyncdb("default", true, true)
}

/* End of file : models.go */
/* Location : ./models/models.go */
