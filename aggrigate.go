package main

import (
	"context"
	"fmt"
)

func handlerAggrigate(s *state, cmd command) error {
	feedURL := "https://www.wagslane.dev/index.xml"

	rss, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return err
	}

	fmt.Printf("RSS feed obj: %v", rss)

	return nil
}
