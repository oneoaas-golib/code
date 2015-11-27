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
	Updated  time.Time `orm:"index;auto_noe;type(datetime)"`
}

//添加初始用户
func AddAdmin() error {
	username := "admin@fun-x.cn"
	password := util.Md5("admin")
	user := &User{
		Username: username,
		Password: password,
		Updated:  time.Now(),
	}
	o := orm.NewOrm()
	created, _, err := o.ReadOrCreate(user, "Username")
	if err == nil {
		if created {
			return nil
		} else {
			return errors.New("用户名已经存在！")
		}
	}
	return err
}

//检测用户名密码正确
func CheckLogin(username, password string) error {
	o := orm.NewOrm()
	user := &User{Username: username}
	err := o.Read(user, "Username")
	if err != nil {
		if err == orm.ErrNoRows {
			return errors.New("用户名不存在！")
		}
		return err
	}
	if util.Md5(password) != user.Password {
		return errors.New("用户密码错误！")
	}
	return nil
}

//添加用户
func AddUser(username, password string) error {
	o := orm.NewOrm()
	user := &User{
		Username: username,
		Password: password,
	}
	created, _, err := o.ReadOrCreate(user, "Username")
	if err == nil {
		if created {
			return nil
		} else {
			return errors.New("用户名已经存在！")
		}
	}
	return err
}

//修改用户
func EditUser(id int64, username, password string) error {
	user := &User{Id: id}
	o := orm.NewOrm()
	err := o.Read(user)
	if err != nil {
		return err
	}
	user.Username = username
	err = o.Read(user, "Username")
	if err == nil && user.Id != id {
		return errors.New("用户名已经存在！")
	}
	user.Password = util.Md5(password)
	_, err = o.Update(user)
	return err
}

//删除用户
func DelUser(id int64) error {
	o := orm.NewOrm()
	user := &User{Id: id}
	_, err := o.Delete(user)
	return err
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
	_, err := o.QueryTable("user").Limit(pagesize).Offset(offset).OrderBy("-Created").All(users)
	return users, err
}

// 获取用户的数量
func GetUserCount() (int64, error) {
	o := orm.NewOrm()
	count, err := o.QueryTable("user").Count()
	return count, err
}

/* End of file 	: user.go */
/* Location 	: ./models/user.go */
