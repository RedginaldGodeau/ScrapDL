package main

import (
	"ScrapDL/internal/scrapper"
	"os"
)

func main() {

	var args = os.Args
	if len(args) < 3 {
		println("[Download path] [Url of website]")
		return
	}

	var path = args[1]
	var url = args[2]

	scrapper.InitScrapper(url, path)

}
