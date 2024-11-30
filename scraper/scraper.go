package scraper

import (
	"context"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/darv86/goserv/shared"
)

type RSSFeed struct {
	Channel struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		Language    string `xml:"language"`
		Item        []struct {
			Title       string `xml:"title"`
			Link        string `xml:"link"`
			Description string `xml:"description"`
			PubDate     string `xml:"pubDate"`
		} `xml:"item"`
	} `xml:"channel"`
}

func Run(scraperConf *shared.ScraperConfig) {
	context := context.Background()
	queries := scraperConf.Queries
	maxFeeds := scraperConf.MaxFeedsAtTime
	ticker := time.NewTicker(time.Second * 5)
	for ; ; <-ticker.C {
		feedsFetched, err := queries.FeedFetchedGetAll(context, int32(maxFeeds))
		if err != nil {
			log.Println(err.Error())
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feedsFetched {
			wg.Add(1)
			log.Println(feed.Name)
			go makeFeedFetch(feed.Url, wg)
			_, err := queries.FeedMarkFetched(context, feed.ID)
			if err != nil {
				log.Println(err.Error())
			}
		}
		wg.Wait()
	}
}

func makeFeedFetch(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Get(url)
	if err != nil {
		log.Println(err.Error())
	}
	body := response.Body
	defer body.Close()
	feedData, err := io.ReadAll(body)
	if err != nil {
		log.Println(err.Error())
	}
	var feed RSSFeed
	err = xml.Unmarshal(feedData, &feed)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(feed.Channel.Title)
}
