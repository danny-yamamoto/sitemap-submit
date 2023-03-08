package main

import (
	"context"
	"fmt"

	"google.golang.org/api/option"
	"google.golang.org/api/webmasters/v3"
)

func main() {
	key := "credentials.json"
	ctx := context.Background()
	webmastersService, err := webmasters.NewService(ctx, option.WithCredentialsFile(key))
	if err != nil {
		fmt.Println("Failed to create Client.")
	}
	siteUrl := "sc-domain:trial-net.co.jp"
	feedpath := "https://www.trial-net.co.jp/sitemap/sitemap-0.xml"
	err = webmastersService.Sitemaps.Submit(siteUrl, feedpath).Do()
	if err != nil {
		fmt.Println("Failed to submit.")
		fmt.Println(err)
	}
}
