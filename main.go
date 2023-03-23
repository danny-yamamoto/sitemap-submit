package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/webmasters/v3"
)

func SitemapSubmit(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	keyJson := os.Getenv("KEY_JSON")
	webmastersService, err := webmasters.NewService(ctx, option.WithCredentialsJSON([]byte(keyJson)))
	if err != nil {
		fmt.Println("Failed to create Client.")
	}
	siteUrl := os.Getenv("SITE_URL")
	fmt.Printf("siteUrl: %v\n", siteUrl)
	feedpath := os.Getenv("FEEDPATH")
	fmt.Printf("feedpath: %v\n", feedpath)
	err = webmastersService.Sitemaps.Submit(siteUrl, feedpath).Do()
	if err != nil {
		fmt.Println("Failed to submit.")
		fmt.Println(err)
	}
	fmt.Println("Updated.")
}

func main() {
	http.HandleFunc("/", SitemapSubmit)
	http.ListenAndServe(":8080", nil)
}
