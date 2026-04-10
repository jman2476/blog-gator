package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jman2476/blog-gator/internal/config"
	"github.com/jman2476/blog-gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	initConfig, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
	}
	var pgrmState state
	pgrmState.config = initConfig

	db, err := sql.Open("postgres", pgrmState.config.DB_url)
	if err != nil {
		fmt.Println(
			fmt.Errorf("Error establishing db connection: %w", err),
		)
		os.Exit(1)
	}

	dbQueries := database.New(db)
	pgrmState.db = dbQueries

	var cmdList = commands{
		make(map[string]func(*state, command) error),
	}
	// cmdList.register("login", handlerLogin)
	// cmdList.register("register", handlerRegister)
	cmdList.registerAll()

	if len(os.Args) < 2 {
		fmt.Println(
			fmt.Errorf("Program expects 2 or more arguments"),
		)
		os.Exit(1)
	}

	var cmd = command{
		os.Args[1],
		os.Args[2:],
	}

	err = cmdList.run(&pgrmState, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (l commands) registerAll() {
	l.register("login", handlerLogin)
	l.register("register", handlerRegister)
	l.register("reset", handlerReset)
	l.register("users", handlerUsers)
	l.register("agg", handlerAggrigate)
	l.register("addfeed", logInMiddleware(handlerAddFeed))
	l.register("feeds", handlerFeeds)
	l.register("follow", logInMiddleware(handlerFollow))
	l.register("following", logInMiddleware(handlerFollowing))
	l.register("unfollow", logInMiddleware(handlerUnfollow))
	l.register("browse", logInMiddleware(handlerBrowse))
}
