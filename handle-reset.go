package main

import (
	"context"
	"errors"
	"fmt"
)

func handleReset(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return errors.New("This command doesn't take any args")
	}

	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return err
	}

	err = s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("All users have been deleted")

	return nil
}
