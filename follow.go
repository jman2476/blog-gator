package main

import "fmt"

func handlerFollow(s *state, cmd command) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf("Need a url for the feed")
	}

}
