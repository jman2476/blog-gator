package main

import "fmt"

type command struct {
	name      string
	arguments []string
}

type commands struct {
	commandList map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	function, ok := c.commandList[cmd.name]
	if !ok {
		return fmt.Errorf("Command %s not found", cmd.name)
	}

	return function(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandList[name] = f
}
