package main

import (
	"fmt"
	"os"

	"github.com/jman2476/blog-gator/internal/config"
)

func main() {
	initConfig, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
	}
	var pgrmState state
	pgrmState.config = initConfig

	var cmdList = commands{
		make(map[string]func(*state, command) error),
	}
	cmdList.register("login", handlerLogin)

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

	// fmt.Printf("%s, %s\n", pgrmState.config.DB_url, pgrmState.config.Username)
}
