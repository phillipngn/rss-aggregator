package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/phillipngn/rss-aggregator/internal/database"
)

type UserVM struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

type FeedVM struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

type FeedFollowVM struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
}

func parseUser(dbUser database.User) UserVM {
	return UserVM{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}

func parseFeed(dbFeed database.Feed) FeedVM {
	return FeedVM{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func parseFeeds(dbFeeds []database.Feed) []FeedVM {
	result := make([]FeedVM, len(dbFeeds))
	for i, feed := range dbFeeds {
		result[i] = parseFeed(feed)
	}
	return result
}

func parseFeedFollow(feedFollow database.FeedFollow) FeedFollowVM {
	return FeedFollowVM{
		ID:        feedFollow.ID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
		FeedID:    feedFollow.FeedID,
		UserID:    feedFollow.UserID,
	}
}

func parseFeedFollows(feedFollows []database.FeedFollow) []FeedFollowVM {
	result := make([]FeedFollowVM, len(feedFollows))
	for i, feedFollow := range feedFollows {
		result[i] = parseFeedFollow(feedFollow)
	}
	return result
}
