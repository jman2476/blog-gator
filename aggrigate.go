package main

import (
	"fmt"
	"time"
)

func handlerAggrigate(s *state, cmd command) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf("Need a duration between requests for aggrigation to prevent spamming servers.")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("Error parsing duration: %w", err)
	}

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

	// return nil
}
