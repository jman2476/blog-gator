package main

import (
	"context"
	"fmt"

	"github.com/jman2476/blog-gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {

	ff, err := s.db.GetFeedFollowsForUser(
		context.Background(),
		user.ID,
	)
	if err != nil {
		return err
	}

	for _, feed := range ff {
		fmt.Println(feed.FeedName)
	}

	return nil
}
