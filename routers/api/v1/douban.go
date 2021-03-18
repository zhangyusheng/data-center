package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/go-echarts/go-echarts/charts"
	"github.com/zhangyusheng/data-center/pkg/logging"
)

func GenDoubanGraph(c *gin.Context) {
	logging.Info(map[string]interface{}{"name":"zhangyusheng"}, "good")
	p := charts.NewPage(orderRouters("豆瓣250")...)
	nameItems := []string{"衣", "食", "住"}
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "my示例图"})
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", []int{10, 20, 30}).
		AddYAxis("商家B", []int{15, 25, 35})


	bar2 := charts.NewBar()
	bar2.SetGlobalOptions(charts.TitleOpts{Title: "my示例图"})
	bar2.AddXAxis(nameItems).
		AddYAxis("商家A", []int{10, 20, 30}).
		AddYAxis("商家B", []int{15, 25, 35})
	p.Add(
		bar,
		bar2,
	)

	p.Render(c.Writer)
}

