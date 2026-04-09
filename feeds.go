package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserByID(
			context.Background(),
			feed.UserID,
		)
		if err != nil {
			return err
		}
		fmt.Printf("Name: %s\nURL: %s\naddedby: %s\n",
			feed.Name.String, feed.Url.String, user.Name.String)
	}

	return nil
}
