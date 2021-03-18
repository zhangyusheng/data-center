package douban

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/zhangyusheng/data-center/common/common_struct"
	"github.com/zhangyusheng/data-center/models"
	"github.com/zhangyusheng/data-center/pkg/logging"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var (
	BaseUrl = "https://movie.douban.com/top250"
)

func GenDoubanData() {
	logInfo := map[string]interface{}{}
	var movies []models.Movie

	pages := GetPages(BaseUrl)
	logInfo["page"] = pages
	for _, page := range pages {
		nUrl := strings.Join([]string{BaseUrl, page.Url}, "")
		logging.Logger.Info(nUrl)
		res := DoubanHttpRequest(nUrl)
		doc, err := goquery.NewDocumentFromReader(res)
		if err != nil {
			log.Println(err)
		}

		movies = append(movies, ParseMovies(doc)...)
	}

	for _, movie := range movies {
		err := models.UpsertMovie(movie)

		if err != nil {
			logging.Logger.WithFields(map[string]interface{}{"movie":movie}).Errorf("insert movie err %s", err.Error())
		}
	}
}

func GetPages(url string) []common_struct.Page {
	res := DoubanHttpRequest(url)

	doc, err := goquery.NewDocumentFromReader(res)
	if err != nil {
		logging.Logger.Error(err.Error())
	}

	return ParsePages(doc)
}

// 分析分页
func ParsePages(doc *goquery.Document) (pages []common_struct.Page) {
	pages = append(pages, common_struct.Page{Page: 1, Url: ""})
	doc.Find("#content > div > div.article > div.paginator > a").Each(func(i int, s *goquery.Selection) {
		page, _ := strconv.Atoi(s.Text())
		url, _ := s.Attr("href")

		pages = append(pages, common_struct.Page{
			Page: page,
			Url:  url,
		})
	})

	return pages
}

// 分析电影数据
func ParseMovies(doc *goquery.Document) (movies []models.Movie) {
	doc.Find("#content > div > div.article > ol > li").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".hd a span").Eq(0).Text()

		subtitle := s.Find(".hd a span").Eq(1).Text()
		subtitle = strings.TrimLeft(subtitle, "  / ")

		other := s.Find(".hd a span").Eq(2).Text()
		other = strings.TrimLeft(other, "  / ")

		desc := strings.TrimSpace(s.Find(".bd p").Eq(0).Text())
		DescInfo := strings.Split(desc, "\n")
		desc = DescInfo[0]

		movieDesc := strings.Split(DescInfo[1], "/")
		year := strings.TrimSpace(movieDesc[0])
		area := strings.TrimSpace(movieDesc[1])
		tag := strings.TrimSpace(movieDesc[2])

		star := s.Find(".bd .star .rating_num").Text()

		comment := strings.TrimSpace(s.Find(".bd .star span").Eq(3).Text())
		compile := regexp.MustCompile("[0-9]")
		comment = strings.Join(compile.FindAllString(comment, -1), "")

		quote := s.Find(".quote .inq").Text()

		movie := models.Movie{
			Title:    title,
			Subtitle: subtitle,
			Other:    other,
			Desc:     desc,
			Year:     year,
			Area:     area,
			Tag:      tag,
			Star:     star,
			Comment:  comment,
			Quote:    quote,
		}

		movies = append(movies, movie)
	})

	return movies
}

func DoubanHttpRequest(url string) io.Reader {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Mobile Safari/537.36")
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")

	//处理返回结果
	response, err := client.Do(request)
	if err != nil {
		logging.Logger.Error(err.Error())
	}
	if response.StatusCode != 200 {
		logging.Logger.Errorf("status code error: %d %s", response.StatusCode, response.Status)
	}
	return response.Body
}
