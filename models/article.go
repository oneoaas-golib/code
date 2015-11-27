package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

//文章结构体
type Article struct {
	Id       int64
	Title    string    `orm:"size(64);unique"`
	Category string    `orm:"size(32)"`
	Content  string    `orm:"size(5000)"`
	Created  time.Time `orm:"index;auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"index;auto_now;type(datetime)"`
	State    int       `orm:"index;default(1)"`
	Views    int64     `orm:"index"`
}

//添加文章
func AddArticle(title, category, content string) error {
	o := orm.NewOrm()
	article := &Article{
		Title:    title,
		Category: category,
		Content:  content,
		Updated:  time.Now(),
		State:    1,
	}
	if created, _, err := o.ReadOrCreate(article, "Title"); err == nil {
		if !created {
			return errors.New("文章标题已经存在")
		}
	} else {
		return err
	}
	//更新分类下面的文章数量
	cate := &Category{Name: category}
	err := o.Read(cate, "Name")
	if err != nil {
		cate.Count++
		_, err = o.Update(cate, "Count")
	}
	return err
}

//修改文章
func EditArticle(id int64, title, category, content string) error {
	o := orm.NewOrm()
	article := &Article{Id: id}
	var oldCategory string
	if o.Read(article) == nil {
		if article.Category != category {
			oldCategory = article.Category
		}
		article.Title = title
		article.Category = category
		article.Content = content
		_, err := o.Update(article)
		if err != nil {
			return err
		}
	}

	if len(oldCategory) > 0 {
		cate := new(Category)
		err := o.QueryTable("category").Filter("name", oldCategory).One(cate)
		if err != nil {
			return err
		}
		cate.Count--
		_, err = o.Update(cate)
	}

	cate := new(Category)
	err := o.QueryTable("category").Filter("name", category).One(cate)
	if err == nil {
		cate.Count++
		_, err = o.Update(cate)
	}
	return err
}

//删除文章
func DelArticle(id int64) error {
	o := orm.NewOrm()
	article := &Article{Id: id}
	var oldCategory string
	err := o.Read(article)
	if err != nil {
		return err
	}
	oldCategory = article.Category
	_, err = o.Delete(article)
	if err != nil {
		return err
	}

	if len(oldCategory) > 0 {
		cate := new(Category)
		err = o.QueryTable("category").Filter("Name", oldCategory).One(cate)
		if err != nil {

		}
		cate.Count--
		_, err = o.Update(cate)
	}
	return err
}

//获取一篇文章
func GetArticle(id int64) (*Article, error) {
	article := &Article{Id: id}
	o := orm.NewOrm()
	if err := o.Read(article); err == nil {
		article.Views++
		_, err = o.Update(article)
		return article, err
	} else {
		return nil, err
	}
}

//获取文章列表
func GetArticles(offset, pagenum int) ([]*Article, error) {
	articles := make([]*Article, 0)
	o := orm.NewOrm()
	_, err := o.QueryTable("article").Filter("State", 1).OrderBy("-Created").Limit(pagenum).Offset(offset).All(&articles)
	return articles, err
}

//获取文章的总数
func GetArticleCount() (count int64, err error) {
	o := orm.NewOrm()
	count, err = o.QueryTable("article").Count()
	return
}

//移动到回收站
func RemoveToTrash(id int64) error {
	o := orm.NewOrm()
	article := &Article{Id: id}
	if err := o.Read(article); err == nil {
		article.State = 0
		_, err = o.Update(article)
		return err
	} else {
		return err
	}
}

//从回收站恢复
func ReturnFromTrash(id int64) error {
	o := orm.NewOrm()
	article := &Article{Id: id}
	if err := o.Read(&article); err == nil {
		article.State = 1
		_, err = o.Update(article)
		return err
	} else {
		return err
	}
}

/* End of file : article.go */
/* Location : ./models/article.go */
