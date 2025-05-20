package main

import (
	"log"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

func ScrapePage(url string, writer Writer, wg *sync.WaitGroup) {
	// Ensure that the WaitGroup is marked as done when the function exits
	// regardless of how it exits (success or error)
	defer wg.Done()

	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	var textContent strings.Builder

	c.OnHTML("p", func(e *colly.HTMLElement) {
		textContent.WriteString(e.Text + "\n")
	})

	c.OnScraped(func(r *colly.Response) {
		err := writer.Write(PageData{
			URL:  url,
			Text: textContent.String(),
		})
		if err != nil {
			log.Printf("Error writing data for %s: %v", url, err)
		} else {
			log.Printf("Scraped %s", url)
		}
	})

	err := c.Visit(url)
	if err != nil {
		log.Printf("Failed to visit %s: %v", url, err)
	}
}
