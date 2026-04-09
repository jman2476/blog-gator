package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jman2476/blog-gator/internal/database"
)

func scrapeFeeds(s *state, cmd command) error {
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

	for _, item := range feed.Channel.Item {
		fmt.Println(item)
		fmt.Println("===========================")
	}

	return nil
}
