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

//初始化函数，注册模型
func init() {
	orm.RegisterModel(new(Category))
}

//设置表名
func (this *Category) TableName() string {
	return "category"
}

//设置引擎
func (this *Category) TableEngine() string {
	return "INNODB"
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
func GetCategories(page string, pagenum int64) ([]*Category, error) {
	p, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	categories := make([]*Category, 0)
	_, err = o.QueryTable("category").Limit(pagenum).Offset((p - 1) * pagenum).All(&categories)
	return categories, err
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	categories := make([]*Category, 0)
	_, err := o.QueryTable("category").All(&categories)
	return categories, err
}

/* End of file : category.go */
/* Location : ./models/category.go */
