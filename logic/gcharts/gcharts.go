package gcharts

import "github.com/go-echarts/go-echarts/charts"

func GenBarGraph(title string, name string, xdata interface{}, ydata interface{}) *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: title})
	bar.AddXAxis(xdata).
		AddYAxis(name, ydata)
	return bar
}

func GenWcGraph(title string, name string, data interface{}) *charts.WordCloud {
	wc := charts.NewWordCloud()
	wc.SetGlobalOptions(charts.TitleOpts{Title: title})
	wc.Add(name, data.(map[string]interface{}), charts.WordCloudOpts{Shape: "circle", SizeRange: []float32{30, 120}})
	return wc
}

func GenLineGraph(title string, name string, xdata interface{}, ydata interface{}) *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: title})
	line.AddXAxis(xdata).AddYAxis(name, ydata, charts.LabelTextOpts{Show: true})
	return line

}
