package scrapper

import (
	"fmt"
	"github.com/cavaliergopher/grab/v3"
	"github.com/gocolly/colly/v2"
	"path"
	"strings"
)

func InitScrapper(baseURL string, storagePath string) {
	var c = colly.NewCollector()
	var urlList []string
	var urlPos = 0

	imageDL(c, path.Join(storagePath, "/images"))
	videoDL(c, path.Join(storagePath, "/videos"))
	linkVisit(c, &urlList, urlList)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnScraped(func(r *colly.Response) {
		if urlPos < len(urlList) {
			err := c.Visit(urlList[urlPos])
			if err != nil {
				fmt.Println("Scrapping error:", err)
			}
		}
		urlPos++
	})

	err := c.Visit(baseURL)
	if err != nil {
		fmt.Println("Scrapping error", err)
		return
	}
}

func videoDL(collector *colly.Collector, path string) {
	collector.OnHTML("video", func(e *colly.HTMLElement) {
		resp, err := grab.Get(path, e.Attr("src"))
		if err != nil {
			fmt.Println("Download Error :", err)
			return
		}
		fmt.Println("Download saved to", resp.Filename)
	})
}

func imageDL(collector *colly.Collector, path string) {
	collector.OnHTML("img", func(e *colly.HTMLElement) {
		resp, err := grab.Get(path, e.Attr("src"))
		if err != nil {
			fmt.Println("Download Error :", err)
			return
		}
		fmt.Println("Download saved to", resp.Filename)
	})
}

func linkVisit(collector *colly.Collector, linksPointer *[]string, links []string) {
	collector.OnHTML("a", func(e *colly.HTMLElement) {
		if strings.Contains(e.Attr("href"), "https://") {
			var alink = e.Request.AbsoluteURL(e.Attr("href"))
			if visited, _ := collector.HasVisited(alink); !visited {
				*linksPointer = append(links, alink)
			}
		}
	})
}
