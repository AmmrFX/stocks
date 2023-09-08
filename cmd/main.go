package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)


func main() {

	stocks := []string{"abou-kir-fertilizers", "al-rajhi-bank"}
	scrapeUrl := "https://sa.investing.com/equities/"

	c := colly.NewCollector(colly.AllowedDomains("sa.investing.com"))

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept-Language", "en-US;q=0.9")
		fmt.Printf("Visiting %s\n", r.URL)
	})

	c.OnHTML(".text-base.font-bold.leading-6.md\\:text-xl.md\\:leading-7.rtl\\:force-ltr", func(e *colly.HTMLElement) {
		data := strings.TrimSpace(e.Text)
		fmt.Println(data)
	})

	for _, stocl := range stocks {
		url := scrapeUrl + stocl
		err := c.Visit(url)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

}
