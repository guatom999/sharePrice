package scrapper

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

func Scrapper(target string) (string, error) {

	var priceText string

	c := colly.NewCollector(
		colly.AllowedDomains("www.instagram.com", "www.facebook.com", "www.settrade.com", "www.tracker.gg", "tracker.gg", "www.set.or.th"),
	)

	url := fmt.Sprintf("https://www.set.or.th/th/market/product/stock/quote/" + target + "/price")

	c.OnHTML("div.site-container div.quote-info", func(e *colly.HTMLElement) {
		priceText = e.ChildText("div.value")
		fmt.Println(priceText)
	})

	// if priceText == "" {
	// 	return "", errors.New("error: price is null")
	// }

	// Visit the URL and start scraping
	err := c.Visit(url)
	if err != nil {
		log.Println("error: cannot get to page:", err)
		return "", err
	}

	return priceText, nil
}
