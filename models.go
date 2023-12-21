package main

import (
	"time"

	"github.com/avearmin/go-blog-aggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Apikey    string    `json:"apikey"`
}

type Feed struct {
	ID            uuid.UUID `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Name          string    `json:"name"`
	Url           string    `json:"apikey"`
	Userid        uuid.UUID `json:"userid"`
	LastFetchedAt time.Time `json:"last_fetched_at"`
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"Title"`
	Url         string    `json:"url"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func dbUserToJSONUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		Apikey:    user.Apikey,
	}
}

func dbFeedToJSONFeed(feed database.Feed) Feed {
	return Feed{
		ID:            feed.ID,
		CreatedAt:     feed.CreatedAt,
		UpdatedAt:     feed.UpdatedAt,
		Name:          feed.Name,
		Url:           feed.Url,
		Userid:        feed.Userid,
		LastFetchedAt: feed.LastFetchedAt.Time,
	}
}

func dbFeedsToJSONFeeds(dbFeeds []database.Feed) []Feed {
	jsonFeeds := make([]Feed, len(dbFeeds))
	for i := range jsonFeeds {
		jsonFeeds[i] = dbFeedToJSONFeed(dbFeeds[i])
	}
	return jsonFeeds
}

func dbFeedFollowToJSONFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.ID,
		FeedID:    dbFeedFollow.FeedID,
		UserID:    dbFeedFollow.UserID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
	}
}

func dbFeedFollowsToJSONFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	jsonFeedFollows := make([]FeedFollow, len(dbFeedFollows))
	for i := range jsonFeedFollows {
		jsonFeedFollows[i] = dbFeedFollowToJSONFeedFollow(dbFeedFollows[i])
	}
	return jsonFeedFollows
}

func dbPostToJSONPost(dbPost database.Post) Post {
	return Post{
		ID:          dbPost.ID,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
		Title:       dbPost.Title,
		Url:         dbPost.Url,
		Description: dbPost.Description.String,
		PublishedAt: dbPost.PublishedAt.Time,
		FeedID:      dbPost.FeedID,
	}
}

func dbPostsToJSONPosts(dbPosts []database.Post) []Post {
	jsonPosts := make([]Post, len(dbPosts))
	for i := range jsonPosts {
		jsonPosts[i] = dbPostToJSONPost(dbPosts[i])
	}
	return jsonPosts
}
