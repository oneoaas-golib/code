package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

//分类的结构体
type Category struct {
	Id          int64
	Name        string `orm:"size(32);unique"`
	Description string `orm:"size(32)"`
	Count       int64
	Created     time.Time `orm:"index;auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"index;auro_now;type(datetime)"`
}

//添加分类
func AddCategory(name, desc string) (err error) {
	o := orm.NewOrm()
	category := &Category{
		Name:        name,
		Description: desc,
		Updated:     time.Now(),
	}
	if created, _, err := o.ReadOrCreate(category, "Name"); err == nil {
		if created {
			return err
		} else {
			err = errors.New("分类名称已经存在！")
		}
	}
	return
}

//修改分类
func EditCategory(id int64, name, desc string) (err error) {
	o := orm.NewOrm()
	cate := &Category{Id: id}
	if err = o.Read(cate); err == nil {
		cate.Name = name
		cate.Description = desc
		_, err = o.Update(cate)
	}
	return
}

//删除分类
func DelCategory(id int64) (err error) {
	o := orm.NewOrm()
	cate := &Category{Id: id}
	if err = o.Read(cate); err == nil {
		if cate.Count > 0 {
			err = errors.New("分类下面还有文章，请先移除文章")
		} else {
			_, err = o.Delete(cate)
		}
	}
	return
}

//获取一个分类
func GetCategory(id int64) (*Category, error) {
	o := orm.NewOrm()
	cate := &Category{Id: id}
	err := o.Read(cate)
	return cate, err
}

//获取分类列表
func GetCategories(offset, pagesize int) ([]*Category, error) {
	o := orm.NewOrm()
	categories := make([]*Category, 0)
	_, err := o.QueryTable("category").Limit(pagesize).Offset(offset).All(&categories)
	return categories, err
}

//获取所有的分类
func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	categories := make([]*Category, 0)
	_, err := o.QueryTable("category").All(&categories)
	return categories, err
}

//获取分类个数
func GetCategoryCount() (count int64, err error) {
	o := orm.NewOrm()
	count, err = o.QueryTable("category").Count()
	return
}

//通过分类名称获取分类数量
func GetCateCountByName(name string) (count int64, err error) {
	o := orm.NewOrm()
	count, err = o.QueryTable("category").Filter("Name", name).Count()
	return
}

/* End of file : category.go */
/* Location : ./models/category.go */
