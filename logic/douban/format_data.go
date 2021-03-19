package douban

import (
	"github.com/go-echarts/go-echarts/charts"
	"github.com/zhangyusheng/data-center/common/util"
	"github.com/zhangyusheng/data-center/logic/gcharts"
	"github.com/zhangyusheng/data-center/models"
	"github.com/zhangyusheng/data-center/pkg/logging"
	"sort"
	"strings"
)

func LoadData() []models.Movie {
	movies, err := models.GetMovies(0,0, map[string]interface{}{})
	if err != nil {
		logging.Logger.Errorf("get movie db error %s", err.Error())
	}
	return movies
}

func GenStarBar(movies []models.Movie) *charts.Bar {
	starData := map[string]int{}
	starIndex := []string{}
	starValue := []int{}
	for _, movie := range movies {
		starData[movie.Star] += 1
		if !util.InArray(movie.Star, starIndex) {
			starIndex = append(starIndex, movie.Star)
		}
	}
	sort.Sort(sort.StringSlice(starIndex))
	for _, idx := range starIndex {
		starValue = append(starValue, starData[idx])
	}

	return gcharts.GenBarGraph("评分分布图","star", starIndex, starValue)
}

func GenWordCloud(movies []models.Movie) *charts.WordCloud {
	tagData := map[string]int{}
	for _, movie := range movies {
		tagStr := movie.Tag
		tagArr := strings.Split(tagStr, " ")
		for _, tag := range tagArr {
			tagData[tag] += 1
		}
	}
	tagDataFin := map[string]interface{}{}
	for k,v := range tagData{
		tagDataFin[k] = v
	}

	logging.Logger.Info(tagData)
    return gcharts.GenWcGraph("高分电影标签分布", "标签分布", tagDataFin)
}

func GenAreaGraph(movies []models.Movie) *charts.Bar {
	areaData := map[string]int{}
	for _, movie := range movies {
		areaStr := movie.Area
		areaArr := strings.Split(areaStr, " ")
		for _, v := range areaArr {
			areaData[v] += 1
		}
	}
	xdata, ydata := []string{}, []int{}
	for k,v := range areaData{
		xdata = append(xdata, k)
		ydata = append(ydata, v)
	}

	return gcharts.GenBarGraph("高分电影地区分布图","地区分布", xdata, ydata)
}