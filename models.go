package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/haroonalbar/rss-aggregator/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

// all this does is convert the dbUser that is in camal case to User which marshals json to snake case
func databaseUsertoUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserId    uuid.UUID `json:"user_id"`
}

func databaseFeedtoFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserId:    dbFeed.UserID,
	}
}

func databaseFeedstoFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, feed := range dbFeeds {
		feeds = append(feeds, databaseFeedtoFeed(feed))
	}
	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserId    uuid.UUID `json:"user_id"`
	FeedId    uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowtoFeedFollow(dbFeedFollows database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollows.ID,
		CreatedAt: dbFeedFollows.CreatedAt,
		UpdatedAt: dbFeedFollows.UpdatedAt,
		UserId:    dbFeedFollows.UserID,
		FeedId:    dbFeedFollows.FeedID,
	}
}

func databaseFeedFollowstoFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}
	for _, feedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, databaseFeedFollowtoFeedFollow(feedFollow))
	}
	return feedFollows
}

// avoided sql.NullSting for Description cause
// in json unmarshelling we will get a nested structure for it
// which is bax ux for to work with insted we use pointer to sting
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

func databasePostToPost(dbPost database.Post) Post {
	var description *string
	// check if description is valid
	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}

	return Post{
		ID:          dbPost.ID,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
		Title:       dbPost.Title,
		Description: description,
		PublishedAt: dbPost.PublishedAt,
		Url:         dbPost.Url,
		FeedID:      dbPost.FeedID,
	}
}

func databasePoststoPosts(dbPosts []database.Post) []Post {
	posts := []Post{}
	for _, post := range dbPosts {
		posts = append(posts, databasePostToPost(post))
	}
	return posts
}
