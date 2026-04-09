package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jman2476/blog-gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("Missing 1 argument: username")
	}

	username := cmd.arguments[0]

	_, err := s.db.GetUser(
		context.Background(), username)
	if err != sql.ErrNoRows {
		fmt.Println(
			fmt.Errorf("Username %s is already in use", username),
		)
		os.Exit(1)
	}

	createArgs := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}
	user, err := s.db.CreateUser(
		context.Background(), createArgs)
	if err != nil {
		return fmt.Errorf("Error registering user: %w", err)
	}

	err = s.config.SetUser(username)
	if err != nil {
		return fmt.Errorf("Error setting user: %w", err)
	}

	fmt.Printf("User %s successfully created\n", username)
	fmt.Println(user)

	return nil
}
