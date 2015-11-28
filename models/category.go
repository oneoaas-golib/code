package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"strconv"
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
func AddCategory(name, desc string) error {
	o := orm.NewOrm()
	category := &Category{
		Name:        name,
		Description: desc,
		Updated:     time.Now(),
	}
	created, _, err := o.ReadOrCreate(category, "Name")
	if err == nil {
		if created {
			return nil
		} else {
			return errors.New("分类名称已经存在！")
		}
	}
	return err
}

//修改分类
func EditCategory(cid, name, desc string) error {
	id, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Category{Id: id}
	err = o.Read(cate)
	if err != nil {
		return err
	}
	cate.Name = name
	cate.Description = desc
	_, err = o.Update(cate)
	return err

}

//删除分类
func DelCategory(cid string) error {
	id, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Category{Id: id}
	err = o.Read(cate)
	if err != nil {
		return err
	}
	if cate.Count > 0 {
		return errors.New("分类下面还有文章，请先移除文章")
	}
	_, err = o.Delete(cate)
	return err
}

//获取一个分类
func GetCategory(cid string) (*Category, error) {
	id, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	cate := &Category{Id: id}
	err = o.Read(cate)
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
func GetCategoryCount() (int64, error) {
	o := orm.NewOrm()
	count, err := o.QueryTable("category").Count()
	return count, err
}

//通过分类名称获取分类数量
func GetCateCountByName(name string) (int64, error) {
	o := orm.NewOrm()
	count, err := o.QueryTable("category").Filter("Name", name).Count()
	return count, err
}

/* End of file : category.go */
/* Location : ./models/category.go */
