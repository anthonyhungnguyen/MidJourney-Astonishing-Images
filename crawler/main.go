package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	resp, err := http.Get(BASE_URL)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// Check if response is OK
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	// Use goquery to fetch the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var AllMetas []ParsedInfo

	doc.Find(TAG).Each(func(i int, s *goquery.Selection) {
		if len(s.Text()) != 0 {
			var data Data
			err := json.Unmarshal([]byte(s.Text()), &data)

			if err != nil {
				log.Fatal(err)
			}

			var imagePaths []string

			for _, job := range data.Props.PageProps.Jobs {
				imagePaths = append(imagePaths, job.ImagePaths...)
			}

			AllMetas = append(AllMetas, ParsedInfo{
				Height:     data.Props.PageProps.Event.Height,
				Width:      data.Props.PageProps.Event.Width,
				ImagePaths: imagePaths,
				Prompt:     data.Props.PageProps.Prompt,
				Type:       data.Props.PageProps.Type,
			})
		}
	})
}
