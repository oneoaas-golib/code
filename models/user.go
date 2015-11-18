package models

import (
	"crypto/md5"
	"encoding/hex"
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
	password := Str2Md5("admin")
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
	if Str2Md5(password) != user.Password {
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
	user.Password = Str2Md5(password)
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

func GetUsers(start, off string) ([]*User, error) {
	limit, err := strconv.ParseInt(start, 10, 64)
	if err != nil {
		return nil, err
	}
	offset, err := strconv.ParseInt(off, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	users := make([]*User, 0)
	_, err = o.QueryTable("user").Limit(limit, offset).OrderBy("-Created").All(users)
	return users, err
}

//md5加密
func Str2Md5(password string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(password))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

/* End of file : user.go */
/* Location : ./models/user.go */
