package main

import (
	"fmt"

	"github.com/jman2476/blog-gator/internal/config"
)

func main() {
	firstConfig, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
	}

	err = firstConfig.SetUser("jeremy")
	if err != nil {
		fmt.Printf("Error setting user: %v\n", err)
	}

	newConfig, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
	}

	fmt.Printf("Config struct: %v\n", newConfig)
}
