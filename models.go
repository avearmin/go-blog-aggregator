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
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"apikey"`
	Userid    uuid.UUID `json:"userid"`
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		Userid:    feed.Userid,
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
