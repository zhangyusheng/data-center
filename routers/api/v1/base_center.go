package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-echarts/go-echarts/charts"
)

const (
	host   = "http://127.0.0.1:8000/api/v1"
	maxNum = 50
)

type router struct {
	name string
	charts.RouterOpts
}

var (
	routers = []router{
		{"豆瓣250", charts.RouterOpts{URL: host + "/GenDoubanGraph", Text: "豆瓣高分电影排行"}},
		{"汇总图表", charts.RouterOpts{URL: host + "/CorePage", Text: "汇总核心图表"}},
	}
)

func orderRouters(chartType string) []charts.RouterOpts {
	for i := 0; i < len(routers); i++ {
		if routers[i].name == chartType {
			routers[i], routers[0] = routers[0], routers[i]
			break
		}
	}

	rs := make([]charts.RouterOpts, 0)
	for i := 0; i < len(routers); i++ {
		rs = append(rs, routers[i].RouterOpts)
	}
	return rs
}

func CorePage(c *gin.Context) {
	p := charts.NewPage(orderRouters("汇总图表")...)
	nameItems := []string{"衣", "食", "住"}
	bar2 := charts.NewBar()
	bar2.SetGlobalOptions(charts.TitleOpts{Title: "my示例图"})
	bar2.AddXAxis(nameItems).
		AddYAxis("商家A", []int{10, 20, 30}).
		AddYAxis("商家B", []int{15, 25, 35})
	p.Add(
		bar2,
	)

	p.Render(c.Writer)
}