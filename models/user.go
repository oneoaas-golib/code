package models

import (
	"code/util"
	"errors"
	"github.com/astaxie/beego/orm"
	"strconv"
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

//设置表名
func (this *User) TableName() string {
	return "user"
}

//设置引擎
func (this *User) TableEngine() string {
	return "INNODB"
}

//注册模型
func init() {
	orm.RegisterModel(new(User))
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

func EditUser(uid, username, password string) error {
	id, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return err
	}
	user := &User{Id: id}
	o := orm.NewOrm()
	err = o.Read(user)
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

func DelUser(uid string) error {
	id, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	user := &User{Id: id}
	_, err = o.Delete(user)
	return err
}

func GetUser(uid string) (*User, error) {
	id, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	user := &User{Id: id}
	err = o.Read(user)
	return user, err
}

func GetUsers(page string, pagenum int64) ([]*User, error) {
	_page, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	users := make([]*User, 0)
	_, err = o.QueryTable("user").Limit(pagenum).Offset((_page - 1) * pagenum).OrderBy("-Created").All(users)
	return users, err
}

/* End of file : user.go */
/* Location : ./models/user.go */
