package main

import (
	"context"
	"fmt"

	"github.com/jman2476/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf("What feed do you want to unfollow?")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("Error getting feed by url: %w", err)
	}

	deleteArgs := database.DeleteFollowForUserParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	err = s.db.DeleteFollowForUser(context.Background(), deleteArgs)
	if err != nil {
		return fmt.Errorf("Error deleting feed from follows: %w", err)
	}

	return nil
}
