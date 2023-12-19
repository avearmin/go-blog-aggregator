package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/avearmin/go-blog-aggregator/internal/database"
)

// What a horrible function; will need to clean this up ALOT
func worker(DB *database.Queries) {
	log.Println("Starting Worker")

	numFeedsToFetch := 10
	interval := 60 * time.Second
	wg := sync.WaitGroup{}
	mu := &sync.Mutex{}

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
		rssToProcess := []RSS{}
		for _, feed := range nextFeeds {
			log.Printf("Worker fetching: %s", feed.Url)
			wg.Add(1)
			go func(feed database.Feed) {
				defer wg.Done()
				rss, err := fetchFromFeed(feed.Url)
				if err != nil {
					log.Printf("Worker encountered error: %s", err)
				}
				mu.Lock()
				log.Printf("Worker adding %s to process queue", feed.Url)
				rssToProcess = append(rssToProcess, rss)
				mu.Unlock()
			}(feed)
		}
		wg.Wait()
		for _, rss := range rssToProcess {
			log.Printf("Worker processed: %s", rss.Channel.Title)
		}
		log.Println("Worker going to sleep")
	}
}
