package models

import (
	"code/util"
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

//用户结构体
type User struct {
	Id       int64
	Username string    `orm:"size(32)"`
	Password string    `orm:"size(64)"`
	Created  time.Time `orm:"index;auto_now_add;type(datetime)"`
	Login    time.Time `orm:"index;auto_noe;type(datetime)"`
	State    int
}

//添加初始用户
func AddAdmin() (err error) {
	user := &User{
		Username: "admin@fun-x.cn",
		Password: util.Md5("admin"),
		Login:    time.Now(),
	}
	o := orm.NewOrm()
	if created, _, err := o.ReadOrCreate(user, "Username"); err == nil {
		if created {
			return err
		} else {
			err = errors.New("用户名已经存在！")
		}
	}
	return
}

//检测用户名密码正确
func CheckLogin(username, password string) (err error) {
	o := orm.NewOrm()
	user := &User{Username: username}
	if err = o.Read(user, "Username"); err == nil {
		if util.Md5(password) != user.Password {
			err = errors.New("用户密码错误！")
			return
		}
	} else if err == orm.ErrNoRows {
		err = errors.New("用户名不存在！")
	}
	return
}

//添加用户
func AddUser(username, password string) (err error) {
	o := orm.NewOrm()
	user := &User{
		Username: username,
		Password: password,
		Login:    time.Now(),
	}
	if created, _, err := o.ReadOrCreate(user, "Username"); err == nil {
		if !created {
			err = errors.New("用户名已经存在！")
			return err
		}
	}
	return err
}

//修改用户
func EditUser(id int64, username, password string) (err error) {
	user := &User{Id: id}
	o := orm.NewOrm()
	if err = o.Read(user); err == nil {
		if user.Id != id {
			err = errors.New("用户名已经存在！")
			return
		}
		user.Username = username
		user.Password = util.Md5(password)
		_, err = o.Update(user)
	}
	return
}

//删除用户
func DelUser(id int64) (err error) {
	o := orm.NewOrm()
	user := &User{Id: id}
	_, err = o.Delete(user)
	return
}

//获取用户
func GetUser(id int64) (*User, error) {
	o := orm.NewOrm()
	user := &User{Id: id}
	err := o.Read(user)
	return user, err
}

//获取用户列表
func GetUsers(offset, pagesize int) ([]*User, error) {
	o := orm.NewOrm()
	users := make([]*User, 0)
	_, err := o.QueryTable("user").Limit(pagesize).Offset(offset).OrderBy("-Created").All(&users)
	return users, err
}

// 获取用户的数量
func GetUserCount() (count int64, err error) {
	o := orm.NewOrm()
	count, err = o.QueryTable("user").Count()
	return
}

/* End of file 	: user.go */
/* Location 	: ./models/user.go */
