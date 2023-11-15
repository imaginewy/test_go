package routers

import (
	_ "gin_blog/docs"
	"gin_blog/middleware"
	"gin_blog/pkg/setting"
	"gin_blog/pkg/upload"
	"gin_blog/routers/api"
	v1 "gin_blog/routers/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouters() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger())
	router.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	gin.SetMode(setting.ServerSetting.RunMode)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/auth", api.GetAuth)
	router.POST("/upload", api.UploadImage)
	apiv1 := router.Group("api/v1")
	apiv1.Use(middleware.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return router
}
