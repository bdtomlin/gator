package main

import (
	"context"
	"errors"
	"fmt"
)

func handleLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Login requires a username")
	}
	name := cmd.args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return err
	}

	fmt.Println("The user has been set")

	return nil
}
