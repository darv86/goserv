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
	interval := scraperConf.TickInterval
	ticker := time.NewTicker(interval)
	for ; ; <-ticker.C {
		feedsFetched, err := queries.FeedFetchedGetAll(context, int32(maxFeeds))
		if err != nil {
			log.Println(err.Error())
		}
		wg := &sync.WaitGroup{}
		for _, feedDb := range feedsFetched {
			wg.Add(1)
			go makeFeedFetch(feedDb, queries, wg)
		}
		wg.Wait()
	}
}

func makeFeedFetch(feedDb database.Feed, queries *database.Queries, wg *sync.WaitGroup) {
	defer wg.Done()
	client := &http.Client{}
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
	if feedData == nil {
		log.Println(err.Error())
	}
	//
	_, err = queries.FeedMarkFetched(context.Background(), feedDb.ID)
	if err != nil {
		log.Println(err.Error())
	}
	//
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
		// log.Println(postDb.ID)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"posts_url_key\"") {
				continue
			}
			log.Println(err.Error())
		}
	}
	log.Printf("The feed %s has %v of posts:", feed.Channel.Title, len(feed.Channel.Item))
}
