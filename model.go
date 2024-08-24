package main

import (
	"time"

	"github.com/KatragaddaLokesh/RSSagg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKeys   string    `json:"api_key"`
}

func databaseUserToUser(dbuser database.User) User {
	return User{
		ID:        dbuser.ID,
		CreatedAt: dbuser.CreatedAt,
		UpdatedAt: dbuser.UpdatedAt,
		Name:      dbuser.Name,
		ApiKeys:   dbuser.ApiKeys,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(dbfeed database.Feed) Feed {
	return Feed{
		ID:        dbfeed.ID,
		CreatedAt: dbfeed.CreatedAt,
		UpdatedAt: dbfeed.UpdatedAt,
		Name:      dbfeed.Name,
		Url:       dbfeed.Url,
		UserID:    dbfeed.UserID,
	}
}

func databaseFeedToFeeds(dbfeeds []database.Feed) []Feed {
	feeds := []Feed{}

	for _, dbfeed := range dbfeeds {
		feeds = append(feeds, databaseFeedToFeed(dbfeed))
	}
	return feeds
}

type FeedsFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowersToFeedFollowers(dbfeed database.FeedsFollow) FeedsFollow {
	return FeedsFollow{
		ID:        dbfeed.ID,
		CreatedAt: dbfeed.CreatedAt,
		UpdatedAt: dbfeed.UpdatedAt,
		UserID:    dbfeed.UserID,
		FeedID:    dbfeed.FeedID,
	}
}

func databaseFeedFollowsToFeedsFollows(dbfeeds []database.FeedsFollow) []FeedsFollow {
	feeds_follows := []FeedsFollow{}

	for _, dbfeed := range dbfeeds {
		feeds_follows = append(feeds_follows, databaseFeedFollowersToFeedFollowers(dbfeed))
	}
	return feeds_follows
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func databasePostToPost(dbfeed database.Post) Post {
	var description *string
	if dbfeed.Description.Valid {
		description = &dbfeed.Description.String
	}
	return Post{
		ID:          dbfeed.ID,
		CreatedAt:   dbfeed.CreatedAt,
		UpdatedAt:   dbfeed.UpdatedAt,
		Title:       dbfeed.Title,
		Description: description,
		PublishedAt: dbfeed.PublishedAt,
		Url:         dbfeed.Url,
		FeedID:      dbfeed.FeedID,
	}
}

func databasePostsToPosts(dbposts []database.Post) []Post {
	posts := []Post{}

	for _, dbpost := range dbposts {
		posts = append(posts, databasePostToPost(dbpost))
	}
	return posts
}
