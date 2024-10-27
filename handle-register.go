package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/bdtomlin/gator/internal/database"
	"github.com/google/uuid"
)

func handleRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("The register command requires a username")
	}
	name := cmd.args[0]

	u, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	})
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(u.Name)
	if err != nil {
		return err
	}

	fmt.Println("The user has been created")
	fmt.Printf("%+v", u)

	return nil
}
