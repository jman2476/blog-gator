package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("Missing 1 argument: username")
	}

	username := cmd.arguments[0]
	err := s.config.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("User has been set to %s\n", username)
	return nil
}
