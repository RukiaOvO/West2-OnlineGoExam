package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var check uint = 1

const (
	username = "root"
	password = "root"
	host     = "127.0.0.1"
	port     = 3306
	Dbname   = "Moviedb"
)

type Moviedata struct {
	gorm.Model

	Title   string
	Score   string
	Img     string
	Comment string
	Pre     string `gorm:"default:NULL"`
}

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect, err = " + err.Error())
	}

	for i := 0; i < 10; i++ {
		db.AutoMigrate(&Moviedata{})

		crawl(db, strconv.Itoa(i*25))
	}
}

func crawl(db *gorm.DB, p string) {
	temp_data := Moviedata{}

	client := http.Client{}
	add := fmt.Sprintf("https://movie.douban.com/top250?start=%s&filter=", p)
	req, err := http.NewRequest("GET", add, nil)
	if err != nil {
		fmt.Println("req err:", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36")

	srep, err := client.Do(req)
	if err != nil {
		fmt.Println("srep err:", err)
	}
	defer srep.Body.Close()

	textdetail, err := goquery.NewDocumentFromReader(srep.Body)
	if err != nil {
		fmt.Println("textdetail failed:", err)
	}

	textdetail.Find("#content > div > div.article > ol > li").
		Each(func(i int, s *goquery.Selection) {
			title := s.Find("div > div.info > div.hd > a > span:nth-child(1)").Text()
			score := s.Find("div > div.info > div.bd > div > span.rating_num").Text()
			img, ok := s.Find("div > div.pic > a > img").Attr("src")
			comment := s.Find("div > div.info > div.bd > div > span:nth-child(4)").Text()
			pre := s.Find("div > div.info > div.bd > p.quote > span").Text()
			if ok {
				temp_data.Title = title
				temp_data.Score = score
				temp_data.Img = img
				temp_data.Comment = comment
				temp_data.Pre = pre
				temp_data.ID = check
			}

			err := db.Create(&temp_data).Error
			if err != nil {
				fmt.Println("Insert err,", err)
				return
			}
			fmt.Println("Insert:", check, "success")
			check++
		})

}
