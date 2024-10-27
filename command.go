package main

import (
	"errors"
	"fmt"
)

type command struct {
	name string
	args []string
}

func newCommand(args []string) (command, error) {
	if len(args) < 2 {
		return command{}, errors.New("A command is required")
	}
	return command{
		name: args[1],
		args: args[2:],
	}, nil
}

type commands struct {
	registered map[string]func(*state, command) error
}

func newCommands() *commands {
	return &commands{
		registered: map[string]func(*state, command) error{},
	}
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registered[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.registered[cmd.name]
	if !ok {
		return fmt.Errorf("Invalid command: %s", cmd.name)
	}

	return f(s, cmd)
}
