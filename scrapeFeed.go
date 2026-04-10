package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jman2476/blog-gator/internal/database"
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
		// fmt.Printf("Link:  %s\n", item.Link)
		fmt.Println("===========================")
	}
}
