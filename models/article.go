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
	Content  string    `orm:"size(10000)"`
	Created  time.Time `orm:"index;auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"index;auto_now;type(datetime)"`
	State    int       `orm:"index;default(1)"`
	Views    int64     `orm:"index"`
}

//添加文章
func AddArticle(title, category, content string) (err error) {
	o := orm.NewOrm()
	article := &Article{
		Title:    title,
		Category: category,
		Content:  content,
		Updated:  time.Now(),
		State:    1,
	}
	if created, _, err = o.ReadOrCreate(article, "Title"); err == nil {
		if created {
			cate := &Category{Name: category}
			err = o.Read(cate, "Name")
			if err == nil {
				cate.Count++
				_, err = o.Update(cate)
			}
		} else {
			err = errors.New("文章标题已经存在")
		}
	}
	return
}

//修改文章
func EditArticle(id int64, title, category, content string) (err error) {
	o := orm.NewOrm()
	article := &Article{Id: id}
	var oldCategory string
	if err = o.Read(article); err == nil {
		if article.Category != category {
			oldCategory = article.Category
		}
		article.Title = title
		article.Category = category
		article.Content = content
		_, err := o.Update(article)
		if err != nil {
			return
		}
	}

	if len(oldCategory) > 0 {
		oldcate := new(Category)
		err = o.QueryTable("category").Filter("Name", oldCategory).One(oldcate)
		if err == nil {
			oldcate.Count--
			_, err = o.Update(oldcate)
			if err != nil {
				return
			}
		}

		newcate := new(Category)
		err = o.QueryTable("category").Filter("Name", category).One(newcate)
		if err == nil {
			newcate.Count++
			_, err = o.Update(newcate)
			if err != nil {
				return
			}
		}
	}
	return
}

//删除文章
func DelArticle(id int64) (err error) {
	o := orm.NewOrm()
	article := &Article{Id: id}
	var oldCategory string
	if err = o.Read(article); err == nil {
		oldCategory = article.Category
		_, err = o.Delete(article)
		if err != nil {
			return
		}
	} else {
		return
	}

	if len(oldCategory) > 0 {
		cate := new(Category)
		err = o.QueryTable("category").Filter("Name", oldCategory).One(cate)
		if err == nil {
			cate.Count--
			_, err = o.Update(cate)
		}
	}
	return
}

//获取一篇文章
func GetArticle(id int64) (article *Article, err error) {
	article.Id = id
	o := orm.NewOrm()
	err = o.Read(article)
	if err == nil {
		article.Views++
		_, err = o.Update(article)
		return
	}
	return
}

//获取文章列表
func GetArticles(offset, pagenum int, state []int) ([]*Article, error) {
	articles := make([]*Article, 0)
	o := orm.NewOrm()
	_, err := o.QueryTable("article").Filter("State__in", state).OrderBy("-Created").Limit(pagenum).Offset(offset).All(&articles)
	return articles, err
}

//获取文章的总数
func GetArticleCount(states []int) (count int64, err error) {
	o := orm.NewOrm()
	count, err = o.QueryTable("article").Filter("State__in", states).Count()
	return
}

//通过分类的名称来获取文章的数量
func GetArticleCountByCate(name string) (count int64, err error) {
	o := orm.NewOrm()
	count, err = o.QueryTable("article").Filter("Category", name).Count()
	return
}

//回收站操作 trash = true 移动到回收站,false 为恢复
func Trash(id int64, trash bool) (err error) {
	o := orm.NewOrm()
	article := &Article{Id: id}
	if err = o.Read(article); err == nil {
		if trash {
			article.State = 0
		} else {
			article.State = 1
		}
		_, err = o.Update(article)
		return
	}
	return
}

/* End of file : article.go */
/* Location : ./models/article.go */
