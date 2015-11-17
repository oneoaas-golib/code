package models

import (
	"time"
)

type User struct {
	Id       int64
	Username string    `orm:"size(32)"`
	Password string    `orm:"size(64)"`
	Created  time.Time `orm:"index;auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"index;auto_noe;type(datetime)"`
}

func (this *User) TableName() string {
	return "user"
}

func AddAdmin() error {
	return nil
}
