package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	ID        int    `gorm:"primary_key" json:"id"`
	Title     string `json:"title"`
	Keyword   string `json:"keyword"`
	Author    string `json:"author"`
	Desc      string `json:"desc"`
	Content   string `json:"content"`
	Numbers   int    `json:"numbers"`
	HeaderPic string `json:"headerPic"`
	Tag       string `gorm:"many2many:Article_tags" json:"tags"`
	Category  string `json:"category"`
	PubTime   string `json:"pubTime"`
	UpTime    string `json:"upTime"`
	Views     int    `json:"views"`
	Likes     int    `json:"likes"`
	Comments  int    `json:"comments"`
}

type Tag struct {
	ID     int    `gorm:"primary_key" json:"id"`
	BlogId int    `json:"blog_id"`
	tag    string `json:"tag""`
}

func CreateArticle(data map[string]interface{}) {
	db.Create(&Article{
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		HeaderPic: data["headerPic"].(string),
		Tag:       data["tag"].(string),
		Category:  data["category"].(string),
	})
}

func GetRecentBlog() (blogs []Article) {
	db.Limit(10).Order("pub_time DESC").Find(&blogs)
	return
}

func GetAllBlogsCount() (count int) {
	var articles []Article
	db.Find(&articles).Count(&count)
	return
}

func GetAllBlog() (blogs []Article) {
	db.Order("pub_time DESC").Find(&blogs)
	return
}

func GetAllBlogs(page, limit int) (blogs []Article) {
	db.Limit(limit).Offset((page - 1) * limit).Find(&blogs)
	return
}

func GetAllCategory() (category []string) {
	var blogs []Article
	db.Select("distinct(category)").Find(&blogs)
	for _, blog := range blogs {
		category = append(category, blog.Category)
	}
	return
}

func GetAllTags() (tags []string) {
	var blogs []Article
	db.Select("distinct(tag)").Find(&blogs)
	for _, blog := range blogs {
		tags = append(tags, blog.Tag)
	}
	return
}

func GetBlogById(id int) (article []Article) {
	db.Where("id = ?", id).Find(&article)
	return
}

func GetBlogsByCategory(pageNumber int, category string) (blogs []Article) {
	PageSize := 8
	db.Where("category = ?", category).Limit(PageSize).Offset((pageNumber - 1) * PageSize).Find(&blogs)
	return
}

func GetBlogsByTag(pageNum int, tag string) (blogs []Article) {
	PageSize := 8
	db.Where("tag = ?", tag).Limit(PageSize).Offset((pageNum - 1) * PageSize).Find(&blogs)
	return
}

func DeleteBlogById(id int) {
	db.Delete(&Article{}, id)
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("pubTime", time.Now().Format("2006-01-02 15:04:05"))
	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("upTime", time.Now().Format("2006-01-02 15:04:05"))
	return nil
}

//CREATE TABLE article(
//	`id` INT UNSIGNED AUTO_INCREMENT,
//	`title` VARCHAR(50) NOT NULL,
//	`author` VARCHAR(20),
//	 `desc`  VARCHAR(200) NOT NULL,
//	`keyword` VARCHAR(50),
//	`header_pic` VARCHAR(5000) NOT NULL,
//	`tag` VARCHAR(20) NOT NULL,
//	`category` VARCHAR(20) NOT NULL,
//	`numbers` INT UNSIGNED,
//	`content`  VARCHAR(10000) NOT NULL,
//	`pub_time` VARCHAR(20),
//	`up_time` VARCHAR(20),
//	`views` INT UNSIGNED,
//	`likes` INT UNSIGNED,
//	`comments` INT UNSIGNED,
//	PRIMARY KEY (`id`)
//)ENGINE=InnoDB DEFAULT CHARSET=utf8;

//create table auth(
//`id` INT UNSIGNED AUTO_INCREMENT,
//`username` VARCHAR(50) NOT NULL,
//`password` VARCHAR(50) NOT NULL,
//PRIMARY KEY (`id`)
//)ENGINE=InnoDB DEFAULT CHARSET=utf8;

//create table class(
//`id` INT UNSIGNED AUTO_INCREMENT,
//`name` VARCHAR(20),
//PRIMARY KEY (`id`)
//)ENGINE=InnoDB DEFAULT CHARSET=utf8;
