package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	//gorm:index，用于声明这个字段为索引，如果你使用了自动迁移功能则会有所影响，在不使用则无影响
	Tag Tag `json:"tag"`
	//Tag字段，实际是一个嵌套的struct，它利用TagID与Tag模型相互关联，
	//在执行查询的时候，能够达到Article、Tag关联查询的功能
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(count)
	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
	//Preload为什么可以查询到每一项关联的Tag
	//PreLoad是一个预加载器，会执行两条sql语句：1.select * from blog_article;
	//2.select * from blog_tag where id IN (1, 2, 3, 4);在查询出结构之后，gorm内部处理对应的映射逻辑
	//将其填充到Article的Tag当中，避免了循环查询
}

func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)
	return
	//Article 如何关联到 Tag ==》gorm本身做了很多约定
	//Article 结构体成员是 TagID 就是外键，gorm会通过类名+ID的方式去找到这两个类的关系
	//Article 结构体成员是 Tag，嵌套在 Article中的Tag结构体，可以通过Related的方式去关联查询
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Update(data)
	return true
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})
	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})
	return true
}
