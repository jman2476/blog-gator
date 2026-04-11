package main

import (
	"github.com/jman2476/gator/internal/config"
	"github.com/jman2476/gator/internal/database"
)

type state struct {
	config *config.Config
	db     *database.Queries
}
