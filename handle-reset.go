package main

import (
	"context"
	"errors"
	"fmt"
	"log"
)

func handleReset(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		log.Fatal(errors.New("This command doesn't take any args"))
	}

	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = s.db.DeleteAllUsers(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("All users have been deleted")

	return nil
}
