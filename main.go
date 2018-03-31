package main

import (
	"fmt"
	"github.com/iikira/BaiduPCS-Go/downloader"
	"github.com/iikira/pixabay-cralwer/pixabay"
	"log"
)

func main() {
	p := pixabay.NewPixabay()
	pis, err := p.GetPhotos(&pixabay.PhotoParameter{
		Lang:      "zh",
		ImageType: "photo",
		Category:  "nature",
		Order:     "popular",
		Page:      2,
		PerPage:   200,
	})
	if err != nil {
		log.Fatalln(err)
	}

	for k := range pis {
		filename := pis[k].Filename()
		fmt.Printf("[%d] %s\n", k, filename)
		der, err := downloader.NewDownloader(pis[k].ImageURL, downloader.Config{
			Client:    p.Client,
			SavePath:  "out/" + filename,
			Parallel:  10,
			CacheSize: 2048,
		})
		if err != nil {
			fmt.Printf("[%d] error: %s\n", k, err)
			continue
		}

		done, err := der.Execute()
		if err != nil {
			fmt.Printf("[%d] error: %s\n", k, err)
			continue
		}

		<-done
	}
}
