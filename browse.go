package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jman2476/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int32 = 2
	if len(cmd.arguments) > 0 {
		userLimit, err := strconv.Atoi(cmd.arguments[0])
		if err != nil {
			return fmt.Errorf("Post limit must be an integer: %w", err)
		}
		limit = int32(userLimit)
	}

	browseArgs := database.GetPostsByUserParams{
		ID:    user.ID,
		Limit: limit,
	}

	posts, err := s.db.GetPostsByUser(context.Background(), browseArgs)
	if err != nil {
		return fmt.Errorf("Error browsing, getting posts: %w", err)
	}

	for _, post := range posts {
		printPost(post)
	}

	return nil
}

func printPost(post database.Post) error {
	fmt.Printf("<<%s>>\n", post.Title)
	fmt.Printf("Pub date: %s\n", post.PublishedAt.Time)
	fmt.Printf("link: %s\n", post.Url)
	fmt.Printf("Text: %s\n", post.Description.String)
	fmt.Println("=====================================")
	return nil
}
