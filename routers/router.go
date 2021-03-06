package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/zhangyusheng/data-center/routers/api"
	"github.com/zhangyusheng/data-center/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/GenDoubanGraph", v1.GenDoubanGraph)
		apiv1.GET("/GenDoubanData", v1.GenDoubanData)
		apiv1.GET("/CorePage", v1.CorePage)
	}

	return r
}
