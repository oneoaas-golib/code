package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

type Category struct {
	Id      int64
	Name    string `orm:"size(32);unique"`
	Count   int64
	Created time.Time `orm:"index;auto_now_add;type(datetime)"`
	Updated time.Time `orm:"index;auro_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Category))
}

func (this *Category) TableName() string {
	return "category"
}

func (this *Category) TableEngine() string {
	return "INNODB"
}

func AddCategory(name string) error {
	o := orm.NewOrm()
	category := &Category{Name: name}
	if created, _, err := o.ReadOrCreate(category, "Name"); err == nil {
		if created {
			return nil
		} else {
			return errors.New("分类名称已经存在！")
		}
	}
	return err
}

func EditCategory(cid, name string) error {
	return nil

}

func DelCategory(cid string) error {
	return nil

}

func GetCategory(cid string) (*Category, error) {
	return nil, nil

}

func GetCategories(start, offset string) ([]*Category, error) {
	return nil, nil

}

/* End of file : category.go */
/* Location : ./models/category.go */
