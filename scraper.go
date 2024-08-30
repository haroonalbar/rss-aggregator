package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/haroonalbar/rss-aggregator/internal/database"
)

// this will be running in the background of the app and should not stop scrapping
func startScrapping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Scraping on %v gorutines on every %s duration ", concurrency, timeBetweenRequest)

	// a ticker is used to trigger an event every set interval.
	ticker := time.NewTicker(timeBetweenRequest)

	// we are using empty parameters for for loop casue
	// it will execute immidiatly after creating a ticker
	// instead of taking the range from ticker channel
	// which will only execute after the fist interval is done
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Printf("Error getting next feeds: %v", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error marking feed as fetched: %v", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Printf("Error getting rss feed from url: %v", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {

		pubTime, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Error parsing date %v : %v", item.PubDate, err)
			continue
		}

		description := sql.NullString{}

		if item.Description != "" {
			description = sql.NullString{String: item.Description, Valid: true}
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Description: description,
			Url:         item.Link,
			PublishedAt: pubTime,
			FeedID:      feed.ID,
		})
		if err != nil {
			// this is expected behaviour if the post already exists it shouldn't create a new one.
			// so loging it is not needed
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Printf("Error creating post: %v", err)
		}

	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
