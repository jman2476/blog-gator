package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jman2476/gator/internal/database"
)

func scrapeFeeds(s *state) error {
	next, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("Can't get next feed: %w", err)
	}

	fetchedArgs := database.MarkFeedFetchedParams{
		UpdatedAt: time.Now(),
		ID:        next.ID,
	}

	err = s.db.MarkFeedFetched(context.Background(), fetchedArgs)

	feed, err := fetchFeed(context.Background(), next.Url)
	if err != nil {
		return fmt.Errorf("Error fetching feed: %w", err)
	}

	fmt.Printf("\n------%s-----\n\n", feed.Channel.Title)
	printFeedTitles(feed.Channel.Item)
	err = savePosts(s, feed.Channel.Item, next.ID)
	if err != nil {
		return err
	}
	// for _, item := range feed.Channel.Item {
	// 	fmt.Println(item)
	// 	fmt.Println("===========================")
	// }

	return nil
}

func printFeedTitles(rssItems []RSSItem) {
	for _, item := range rssItems {
		fmt.Printf("Title: %s\n", item.Title)
		fmt.Printf("Date:  %s\n", item.PubDate)
		fmt.Printf("Link:  %s\n", item.Link)
		fmt.Println("===========================")
	}
}

func savePosts(s *state, rssItems []RSSItem, feedID uuid.UUID) error {

	for _, item := range rssItems {
		description_exists := item.Description != ""
		pubDate_exists := item.PubDate != ""
		pubDate := timeStringToStamp(item.PubDate)

		postArgs := database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     item.Title,
			Url:       item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  description_exists,
			},
			PublishedAt: sql.NullTime{
				Time:  pubDate,
				Valid: pubDate_exists,
			},
			FeedID: feedID,
		}
		_, err := s.db.CreatePost(context.Background(), postArgs)
		if err != nil {
			if strings.Contains(err.Error(), "posts_url_key") {
				continue
			}
			return fmt.Errorf("Error creating post: %w", err)
		}
	}

	return nil
}

func timeStringToStamp(tString string) time.Time {
	formats := []string{
		"ANSIC",
		"UnixDate",
		"RubyDate",
		"RFC822",
		"RFC822Z",
		"RFC850",
		"RFC1123",
		"RFC1123Z",
		"RFC3339",
		"RFC3339Nano",
		"Kitchen",
		"Stamp",
		"StampMilli",
		"StampMicro",
		"StampNano",
		"DateTime",
		"DateOnly",
		"TimeOnly",
	}

	for _, format := range formats {
		timeStamp, err := time.Parse(format, tString)
		if err == nil {
			return timeStamp
		}
	}

	return time.Time{}
}
