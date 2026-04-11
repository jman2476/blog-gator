package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jman2476/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf("Need a url for the feed")
	}

	feed, err := s.db.GetFeedByURL(
		context.Background(), cmd.arguments[0],
	)
	if err != nil {
		return nil
	}

	feedfollowArgs := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	ff, err := s.db.CreateFeedFollow(
		context.Background(),
		feedfollowArgs,
	)
	if err != nil {
		return err
	}

	// fmt.Printf("%s :: %s", ff.UserName, ff.FeedName)
	printFeedFollow(ff)
	return nil
}

func printFeedFollow(ff database.CreateFeedFollowRow) {
	fmt.Printf("User:		%s\n", ff.UserName)
	fmt.Printf("Feed:		%s\n", ff.FeedName)
}
