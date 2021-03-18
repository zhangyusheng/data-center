package douban

import (
	"github.com/go-echarts/go-echarts/charts"
	"github.com/zhangyusheng/data-center/common/util"
	"github.com/zhangyusheng/data-center/models"
	"github.com/zhangyusheng/data-center/pkg/logging"
	"sort"
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

	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "评分分布图"})
	bar.AddXAxis(starIndex).
		AddYAxis("star", starValue)
	return bar
}