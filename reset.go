package main

import (
	"context"
	"fmt"
	"os"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.ClearUsers(context.Background())
	if err != nil {
		fmt.Println("Error Users table not reset")
		os.Exit(1)
	}
	fmt.Println("Users table successfully cleared")
	os.Exit(0)
	return nil
}
