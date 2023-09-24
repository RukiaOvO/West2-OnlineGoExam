package main

import (
	"fmt"
	"fzu_crawl/data"
	model2 "fzu_crawl/data/model"
	"github.com/PuerkitoBio/goquery"
	"gorm.io/gorm"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

const (
	headerSetting = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36"
)

var count int = 1

func main() {
	startPage := 155

	db := model.InitDB()
	sTime := time.Now()
	for {
		err := crawl(db, strconv.Itoa(startPage))
		if err != nil {
			panic(err)
		}
		startPage++
		if startPage > 278 {
			break
		}
	}
	eTime := time.Now()

	fmt.Printf("总共爬取 %d 页, 耗时 %s\n", count, eTime.Sub(sTime))
}

func crawl(db *gorm.DB, p string) error {
	homeUrl := "https://info22.fzu.edu.cn/lm_list.jsp?totalpage=956&PAGENUM=" + p + "&wbtreeid=1460"
	homeClient := http.Client{}
	homeReq, err := http.NewRequest("GET", homeUrl, nil)
	if err != nil {
		return err
	}
	homeReq.Header.Set("User-Agent", headerSetting)

	homeRes, err := homeClient.Do(homeReq)
	if err != nil {
		return err
	}
	defer homeRes.Body.Close()

	homeDetail, err := goquery.NewDocumentFromReader(homeRes.Body)
	if err != nil {
		return err
	}

	for i := 1; i <= 20; i++ {
		tempSelect := fmt.Sprintf("body > div.sy-content > div > div.right.fr > div.list.fl > ul > li:nth-child(%d)", i)
		date := homeDetail.Find(tempSelect).Find("p > span").Text()
		if !dateCheck(date) {
			continue
		}

		author := homeDetail.Find(tempSelect).Find("p > a.lm_a").Text()
		author = author[3 : len(author)-3]
		title, _ := homeDetail.Find(tempSelect).Find("p > a:nth-child(2)").Attr("title")
		link, _ := homeDetail.Find(tempSelect).Find("p > a:nth-child(2)").Attr("href")
		text := "https://info22.fzu.edu.cn/" + link
		id := pickId(text)
		innerUrl := fmt.Sprintf("https://info22.fzu.edu.cn/system/resource/code/news/click/dynclicks.jsp?clickid=%d&owner=1768654345&clicktype=wbnews", id)

		innerClient := http.Client{}
		innerReq, err := http.NewRequest("GET", innerUrl, nil)
		if err != nil {
			return err
		}
		innerReq.Header.Set("User-Agent", headerSetting)
		innerRes, err := innerClient.Do(innerReq)
		if err != nil {
			return err
		}
		tempNums, err := io.ReadAll(innerRes.Body)
		if err != nil {
			return err
		}
		nums := string(tempNums)

		newData := model2.News{}
		newData.Title = title
		newData.Author = author
		newData.Date = date
		newData.Text = text
		newData.Nums = nums

		db.Create(&newData)
		fmt.Println("Insert data", count)
		count++
	}
	return nil
}

func dateCheck(d string) bool {
	numCheck := regexp.MustCompile("[0-9]+")
	numList := numCheck.FindAllString(d, -1)
	tempS := numList[0] + numList[1] + numList[2]
	tempI, err := strconv.Atoi(tempS)
	if err != nil {
		panic(err)
	}
	if 20200101 <= tempI && tempI <= 20210901 {
		return true
	}
	return false
}

func pickId(s string) int {
	idCheck := regexp.MustCompile("[0-9]+")
	sId := idCheck.FindAllString(s, -1)
	Id, err := strconv.Atoi(sId[2])
	if err != nil {
		panic(err)
	}
	return Id
}
