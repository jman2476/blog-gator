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

	var min time.Duration = 5000000000
	if timeBetweenRequests < min {
		return fmt.Errorf("Cannot set time_between_requests to less than 5 seconds.")
	}

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			return err
		}
	}

	// return nil
}
