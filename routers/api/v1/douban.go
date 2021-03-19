package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangyusheng/data-center/pkg/app"
	"github.com/zhangyusheng/data-center/pkg/e"
	"net/http"

	"github.com/go-echarts/go-echarts/charts"
	"github.com/zhangyusheng/data-center/logic/douban"
	"github.com/zhangyusheng/data-center/pkg/logging"
)

func GenDoubanGraph(c *gin.Context) {
	logging.Logger.Info("start gen douban graph")
	p := charts.NewPage(orderRouters("豆瓣250")...)
	movies := douban.LoadData()

	starBar := douban.GenStarBar(movies)
	tagWc := douban.GenWordCloud(movies)
	areaBar := douban.GenAreaGraph(movies)
	p.Add(
		starBar,
		tagWc,
		areaBar,
	)

	p.Render(c.Writer)
}

func GenDoubanData(c *gin.Context) {
	appG := app.Gin{c}
	logging.Logger.WithFields(map[string]interface{}{}).Info("start parsing douban")
	douban.GenDoubanData()
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

