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

	nameItems := []string{"衣", "食", "住"}
	bar2 := charts.NewBar()
	bar2.SetGlobalOptions(charts.TitleOpts{Title: "my示例图"})
	bar2.AddXAxis(nameItems).
		AddYAxis("商家A", []int{10, 20, 30}).
		AddYAxis("商家B", []int{15, 25, 35})
	p.Add(
		starBar,
		bar2,
	)

	p.Render(c.Writer)
}

func GenDoubanData(c *gin.Context) {
	appG := app.Gin{c}
	logging.Logger.WithFields(map[string]interface{}{}).Info("start parsing douban")
	douban.GenDoubanData()
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

