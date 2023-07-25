package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/huason/go-spider/model"
	"github.com/huason/go-spider/parse"
	"log"
	"net/http"
	"strings"
)

var (
	BaseUrl = "https://movie.douban.com/top250"
)

// Add 新增数据
func Add(movies []parse.DoubanMovie) {
	for index, movie := range movies {
		if err := model.DB.Create(&movie).Error; err != nil {
			log.Printf("db.Create index: %d, err : %v", index, err)
		}
	}
}

// Start 开始爬取
func Start() {
	var movies []parse.DoubanMovie
	client := &http.Client{}

	pages := parse.GetPages(BaseUrl)
	for _, page := range pages {
		url := strings.Join([]string{BaseUrl, page.Url}, "")
		request, err := http.NewRequest("GET", url, nil)
		//设置header属性
		request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36 OPR/66.0.3515.115")
		if err != nil {
			fmt.Println(err)
		}
		response, _ := client.Do(request)
		doc, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			log.Println(err)
		}
		movies = append(movies, parse.ParseMovies(doc)...)
	}
	Add(movies)
	fmt.Println("脚本执行完成")
}

func main() {
	Start()
	defer model.CloseDB()
}
