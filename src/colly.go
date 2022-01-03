package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

func collydemo() {

	c := colly.NewCollector()

	// 抓標籤Tag
	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println("1", e.Text)
	})

	// 抓屬性值 AttrVal
	c.OnHTML("meta[name='description']", func(e *colly.HTMLElement) {
		fmt.Println("2", e.Attr("content")) // 抓此Tag中的name屬性 來找出此Tag，再印此Tag中的content屬性
	})

	// 抓類別Class 名稱
	c.OnHTML(".qa-list__title--ironman", func(e *colly.HTMLElement) {
		fmt.Println("3", e.Text)
	})

	// 抓唯一識別 ID
	c.OnHTML("#read_more", func(e *colly.HTMLElement) {
		fmt.Println("4", e.Text)
	})

	c.OnRequest(func(r *colly.Request) { // iT邦幫忙需要寫這一段 User-Agent才給爬
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	c.Visit("https://ithelp.ithome.com.tw/users/20125192/ironman/3155")
}

func collydemo2() {

	c := colly.NewCollector()

	c.OnHTML(".qa-list__title-link", func(e *colly.HTMLElement) {
		// fmt.Println(e.Text, e.Attr("href"))
		// e.Text 印出 <a> tag 裡面的文字，也就是文章標題
		// e.Attr("href") 則是找到 <a> tag裡面的 href元素

		linksStr := e.Attr("href")
		fmt.Println("linksStr = ", linksStr)
		linksStr = strings.Replace(linksStr, " ", "", -1)  // 把空白以""取代掉
		linksStr = strings.Replace(linksStr, "\n", "", -1) // 把空白以""取代掉
		links := strings.Split(linksStr, "\n")             // 以換行符號(\n)做為分隔來切割字串
		// fmt.Println("links = ", links)
		for _, link := range links {
			fmt.Println("link = ", link)
			c2 := colly.NewCollector() // 這邊要在迴圈一開始再宣告一個 Collector，才不會與原本的混合到
			c2.OnHTML(".qa-markdown", func(e2 *colly.HTMLElement) {
				fmt.Println(e2.Text) // 印出 qa-markdown class中的文字（Go繁不及備載 文章的內文）
			})

			// c.OnResponse(func(r *colly.Response) {
			// 	fmt.Println(string(r.Body)) // 印出所有返回的Response物件r.Body
			// })

			c2.OnRequest(func(r *colly.Request) { // iT邦幫忙需要寫這一段 User-Agent才給爬
				r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
			})
			c2.Visit(link) // 找到<a>連結網址後，點進去訪問
		}
	})

	c.OnRequest(func(r *colly.Request) { // iT邦幫忙需要寫這一段 User-Agent才給爬
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	c.Visit("https://ithelp.ithome.com.tw/users/20125192/ironman/3155?page=1")
	c.Visit("https://ithelp.ithome.com.tw/users/20125192/ironman/3155?page=2")
	c.Visit("https://ithelp.ithome.com.tw/users/20125192/ironman/3155?page=3")
}
