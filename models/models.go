package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//数据库的链接
const DB_NAME = "root:root@/beego?charset=utf8&loc=Local"

//注册模型
func init() {
	orm.RegisterModel(new(Article), new(User), new(Category))
}

//初始化数据库
func RegisterDB() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", DB_NAME, 30)
	orm.RunSyncdb("default", false, true)
}

/* End of file : models.go */
/* Location : ./models/models.go */
