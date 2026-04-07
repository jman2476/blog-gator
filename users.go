package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	current := s.config.Username

	for _, user := range users {
		str := fmt.Sprint(user.Name.String)
		if str == current {
			str += " (current)"
		}
		fmt.Println(str)
	}

	return nil
}
