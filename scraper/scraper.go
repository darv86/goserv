package scraper

import (
	"context"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/darv86/goserv/internal/database"
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
		for _, feedDb := range feedsFetched {
			wg.Add(1)
			go makeFeedFetch(feedDb, queries, wg)
			_, err := queries.FeedMarkFetched(context, feedDb.ID)
			if err != nil {
				log.Println(err.Error())
			}
		}
		wg.Wait()
	}
}

func makeFeedFetch(feedDb database.Feed, queries *database.Queries, wg *sync.WaitGroup) {
	defer wg.Done()
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Get(feedDb.Url)
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
	for _, post := range feed.Channel.Item {
		_, err := queries.PostCreate(context.Background(), database.PostCreateParams{
			Title:  post.Title,
			Url:    post.Link,
			FeedID: feedDb.ID,
		})
		if strings.Contains(err.Error(), "duplicate key") {
			continue
		}
		if err != nil {
			log.Println(err.Error())
		}
	}
	log.Printf("The feed %s has %v of posts:", feed.Channel.Title, len(feed.Channel.Item))
}
