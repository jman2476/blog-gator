package main

import (
	"github.com/jman2476/blog-gator/internal/config"
	"github.com/jman2476/blog-gator/internal/database"
)

type state struct {
	config *config.Config
	db     *database.Queries
}
