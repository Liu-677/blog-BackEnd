package routers

import (
	"Blog_Backend/middleware"
	"Blog_Backend/middleware/jwt"
	"Blog_Backend/pkg/setting"
	"Blog_Backend/routers/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)

	apiBlog := r.Group("/api/blogs")
	{
		apiBlog.GET("/getAllCategory", api.GetAllCategory)
		apiBlog.GET("/getAllTags", api.GetAllTags)
		apiBlog.GET("/getBlogById", api.GetBlogById)
		apiBlog.GET("/getRecentBlog", api.GetRecentBlog)
		apiBlog.GET("/getAllBlogsCount", api.GetAllBlogsCount)
		apiBlog.GET("/getBlogsByCategory", api.GetBlogsByCategory)
		apiBlog.GET("/getAllBlogs", api.GetAllBlogs)
		apiBlog.GET("/getBlogsByTag", api.GetBlogsByTag)
		apiBlog.GET("/GetBlogsByPage", api.GetBlogsByPage)
	}

	apiAdmin := r.Group("/api/admin")
	apiAdmin.Use(jwt.JWT())
	{
		apiAdmin.GET("/getAllBlogsCount", api.GetAllBlogsCount)
		apiAdmin.GET("/getAllBlogs", api.GetAllBlogs)
		apiAdmin.GET("/GetBlogsByPage", api.GetBlogsByPage)
		apiAdmin.GET("/getAllCategory", api.GetAllCategory)
		apiAdmin.DELETE("/DeleteBlogById", api.DeleteBlogById)
		apiAdmin.POST("/createBlog", api.CreateArticle)
		apiAdmin.POST("/UpLoadImg", api.UpLoadImg)
	}
	return r
}
