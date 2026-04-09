package main

import (
	"context"
	"database/sql"
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
		context.Background(),
		sql.NullString{
			String: s.config.Username,
			Valid:  true,
		})
	if err != nil {
		return err
	}

	feedArgs := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: sql.NullString{
			String: cmd.arguments[0],
			Valid:  true,
		},
		Url: sql.NullString{
			String: cmd.arguments[1],
			Valid:  true,
		},
		UserID: user.ID,
	}

	feed, err := s.db.CreateFeed(
		context.Background(),
		feedArgs,
	)
	if err != nil {
		return err
	}

	fmt.Println(feed)
	return nil
}
