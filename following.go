package main

import (
	"context"
	"database/sql"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
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

	ff, err := s.db.GetFeedFollowsForUser(
		context.Background(),
		user.ID,
	)

	for _, feed := range ff {
		fmt.Println(feed.FeedName.String)
	}

	return nil
}
