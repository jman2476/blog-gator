package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("Missing 1 argument: username")
	}

	username := cmd.arguments[0]

	user, err := s.db.GetUser(
		context.Background(), username)
	if err == sql.ErrNoRows {
		fmt.Println(
			fmt.Errorf("Username %s not in database", username),
		)
		os.Exit(1)
	}

	err = s.config.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("User has been set to %s\n", username)
	printUser(user)

	return nil
}
