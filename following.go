package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	user, err := s.db.GetUser(
		context.Background(), s.config.Username)
	if err != nil {
		return err
	}

	ff, err := s.db.GetFeedFollowsForUser(
		context.Background(),
		user.ID,
	)

	for _, feed := range ff {
		fmt.Println(feed.FeedName)
	}

	return nil
}
