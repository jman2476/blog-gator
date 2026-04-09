package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jman2476/blog-gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.arguments) < 2 {
		return fmt.Errorf("Need 2 arguments: [title] [url]\n\rHave %d argument(s)", len(cmd.arguments))
	}

	user, err := s.db.GetUser(
		context.Background(), s.config.Username)
	if err != nil {
		return err
	}

	feedArgs := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.arguments[0],
		Url:       cmd.arguments[1],
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(
		context.Background(),
		feedArgs,
	)
	if err != nil {
		return err
	}

	feedfollowArgs := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	_, err = s.db.CreateFeedFollow(
		context.Background(),
		feedfollowArgs,
	)

	printFeed(feed)
	return nil
}

func printFeed(f database.Feed) {
	fmt.Printf("ID:			  %s\n", f.ID)
	fmt.Printf("Created At:   %s\n", f.CreatedAt)
	fmt.Printf("Updated At:   %s\n", f.UpdatedAt)
	fmt.Printf("Name: 		  %s\n", f.Name)
	fmt.Printf("URL: 		  %s\n", f.Url)
	fmt.Printf("User ID: 	  %s\n", f.UserID)
}
