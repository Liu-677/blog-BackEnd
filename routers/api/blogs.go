package api

import (
	"Blog_Backend/models"
	"Blog_Backend/pkg/e"
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
	"os"
	"regexp"
)

func GetRecentBlog(ctx *gin.Context) {
	recentBlogs := models.GetRecentBlog()
	code := e.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"blogs": recentBlogs,
	})
}

func GetAllBlogsCount(ctx *gin.Context) {
	count := models.GetAllBlogsCount()
	valid := validation.Validation{}
	valid.Min(count, 1, "count").Message("ID必须大于0")
	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
	} else {
		code = e.ERROR_NOT_EXIST_ARTICLE
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"total": count,
	})
}

func GetAllBlogs(ctx *gin.Context) {
	blogs := models.GetAllBlog()
	code := e.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"blogs": blogs,
	})
}

func GetAllCategory(ctx *gin.Context) {
	categorys := models.GetAllCategory()
	code := e.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": categorys,
	})
}

func GetAllTags(ctx *gin.Context) {
	tags := models.GetAllTags()
	code := e.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"info": tags,
	})
}

func GetBlogById(ctx *gin.Context) {
	id := com.StrTo(ctx.Query("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id 必须大于0")
	var blogs []models.Article
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		blogs = models.GetBlogById(id)
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"blog": blogs,
	})
}

func GetBlogsByCategory(ctx *gin.Context) {
	page := com.StrTo(ctx.Query("pageNum")).MustInt()
	category := ctx.Query("category")
	valid := validation.Validation{}
	valid.Min(page, 1, "page").Message("pageNum 必须大于0")
	code := e.INVALID_PARAMS
	var blogs []models.Article
	if !valid.HasErrors() {
		code = e.SUCCESS
		blogs = models.GetBlogsByCategory(page, category)
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"data":  blogs,
		"total": len(blogs),
	})
}

func GetBlogsByTag(ctx *gin.Context) {
	page := com.StrTo(ctx.Query("pageNum")).MustInt()
	tag := ctx.Query("tag")
	valid := validation.Validation{}
	valid.Min(page, 1, "page").Message("pageNum 必须大于0")
	code := e.INVALID_PARAMS
	var blogs []models.Article
	if !valid.HasErrors() {
		code = e.SUCCESS
		blogs = models.GetBlogsByTag(page, tag)
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"data":  blogs,
		"total": len(blogs),
	})
}

func GetBlogsByPage(ctx *gin.Context) {
	page := com.StrTo(ctx.Query("page")).MustInt()
	limit := com.StrTo(ctx.Query("limit")).MustInt()
	valid := validation.Validation{}
	valid.Min(page, 1, "page").Message("page 必须大于0")
	valid.Min(limit, 1, "limit").Message("limit 必须大于0")
	code := e.INVALID_PARAMS
	var blogs []models.Article
	if !valid.HasErrors() {
		code = e.SUCCESS
		blogs = models.GetAllBlogs(page, limit)
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"blogs": blogs,
	})
}

func DeleteBlogById(ctx *gin.Context) {
	id := com.StrTo(ctx.Query("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "page").Message("pageNum 必须大于0")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		models.DeleteBlogById(id)
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

func CreateArticle(ctx *gin.Context) {
	data := make(map[string]interface{}, 0)
	data["title"] = ctx.PostForm("title")
	data["category"] = ctx.PostForm("category")
	data["headerPic"] = ctx.PostForm("headerPic")
	data["tag"] = ctx.PostForm("tag")
	data["content"] = ctx.PostForm("content")
	data["desc"] = ctx.PostForm("desc")

	code := e.INVALID_PARAMS
	fmt.Println(data)
	if data["title"] != nil && data["category"] != nil && data["tag"] != nil && data["content"] != nil && data["desc"] != nil {
		models.CreateArticle(data)
		code = e.SUCCESS
	} else {
		fmt.Println("有缺失值,创建失败")
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

func UpLoadImg(ctx *gin.Context) {
	hostPath := "47.98.112.59:6770/"
	filename := ctx.PostForm("name")

	base64_image_content := ctx.PostForm("img")

	code := e.INVALID_PARAMS

	if err := WriteFile("/home/imgs", filename, base64_image_content); err {

	} else {
		code = e.SUCCESS
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"url":  hostPath + filename,
	})
}

func WriteFile(path string, filename string, base64_image_content string) bool {

	b, _ := regexp.MatchString(`^data:\s*image\/(\w+);base64,`, base64_image_content)
	if !b {
		return false
	}

	re, _ := regexp.Compile(`^data:\s*image\/(\w+);base64,`)
	base64Str := re.ReplaceAllString(base64_image_content, "")
	file := path + "/" + filename
	byte, _ := base64.StdEncoding.DecodeString(base64Str)

	//err := ioutil.WriteFile(file, byte, 0666)
	err := os.WriteFile(file, byte, 0666)
	if err != nil {
		log.Println(err)
	}
	return false
}
