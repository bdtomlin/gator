package main

import (
	"context"
	"errors"
	"fmt"
)

func handleUsers(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return errors.New("This command doesn't take any args")
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	if len(users) == 0 {
		fmt.Println("There are no users")
	}

	for _, user := range users {
		printUserName(s, user.Name)
	}
	return nil
}

func printUserName(s *state, name string) {
	if name == s.cfg.CurrentUserName {
		fmt.Printf("* %s (current)\n", name)
	} else {
		fmt.Printf("* %s\n", name)
	}
}
