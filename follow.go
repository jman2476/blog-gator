package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jman2476/blog-gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf("Need a url for the feed")
	}

	user, err := s.db.GetUser(
		context.Background(),
		sql.NullString{
			String: s.config.Username,
			Valid:  true,
		},
	)
	if err != nil {
		return err
	}

	feed, err := s.db.GetFeedByURL(
		context.Background(),
		sql.NullString{
			String: cmd.arguments[0],
			Valid:  true,
		},
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

	fmt.Printf("%s :: %s", ff.UserName.String, ff.FeedName.String)
	return nil
}
