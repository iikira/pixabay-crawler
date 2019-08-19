package main

import (
	"flag"
	"fmt"
	"github.com/iikira/BaiduPCS-Go/requester/downloader"
	"github.com/iikira/pixabay-crawler/pixabay"
	"log"
	"os"
)

var (
	page int
)

func init() {
	flag.IntVar(&page, "p", 1, "page of photo")
	flag.Parse()
}

func main() {
	p := pixabay.NewPixabay()
	pis, err := p.GetPhotos(&pixabay.PhotoParameter{
		Lang:      "zh",
		ImageType: "photo",
		Category:  "nature",
		Order:     "popular",
		Page:      page,
		PerPage:   200,
	})
	if err != nil {
		log.Fatalln(err)
	}

	for k := range pis {
		filename := pis[k].Filename()
		fmt.Printf("[%d] %s\n", k, filename)

		savePath := "out/" + filename
		_, err := os.Stat(savePath)
		if err == nil {
			fmt.Printf("[%d] 已存在\n", k)
			continue
		}

		if os.IsExist(err) {
			fmt.Printf("[%d] error: %s\n", k, err)
			continue
		}

		file, err := os.OpenFile(savePath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("[%d] error: %s\n", k, err)
			continue
		}

		der := downloader.NewDownloader(pis[k].ImageURL, file, &downloader.Config{
			MaxParallel: 1,
			CacheSize:   20480,
		})

		err = der.Execute()
		if err != nil {
			fmt.Printf("[%d] error: %s\n", k, err)
			continue
		}
	}
}
