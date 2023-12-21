package main

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/avearmin/go-blog-aggregator/internal/database"
	"github.com/google/uuid"
)

// What a horrible function; will need to clean this up ALOT
func worker(DB *database.Queries) {
	log.Println("Starting Worker")

	numFeedsToFetch := 10
	interval := 60 * time.Second
	wg := &sync.WaitGroup{}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		<-ticker.C
		log.Println("Worker waking up")
		nextFeeds, err := DB.GetNextFeedsToFetch(context.TODO(), int32(numFeedsToFetch))
		if err != nil {
			log.Printf("Worked encountered error: %s", err)
			continue
		}

		for _, feed := range nextFeeds {
			log.Printf("Worker fetching: %s", feed.Url)
			wg.Add(1)
			go func(wg *sync.WaitGroup, DB *database.Queries, feed database.Feed) {
				defer wg.Done()
				DB.MarkFeedFetched(context.Background(), feed.ID)
				rss, err := fetchFromFeed(feed.Url)
				if err != nil {
					log.Printf("Worker encountered error: %s", err)
				}

				for _, post := range rss.Channel.Item {
					var publishedAt sql.NullTime
					pubDate, err := time.Parse(time.RFC1123, post.PubDate)
					if err != nil {
						log.Println("Worker encountered error parsing pub date from post. Will insert NULL.")
						publishedAt = sql.NullTime{Valid: false}
					} else {
						publishedAt = sql.NullTime{Time: pubDate, Valid: true}
					}
					_, err = DB.CreatePost(context.TODO(), database.CreatePostParams{
						ID:          uuid.New(),
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
						Title:       post.Title,
						Url:         post.Link,
						Description: sql.NullString{String: post.Description, Valid: true},
						PublishedAt: publishedAt,
						FeedID:      feed.ID,
					})
					if err != nil {
						return
					}
					log.Printf("Inserted %s into DB", post.Title)
				}
			}(wg, DB, feed)
		}
		wg.Wait()
		log.Println("Worker going to sleep")
	}
}
