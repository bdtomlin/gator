package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/bdtomlin/gator/internal/database"
)

func handleUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return errors.New("This command requires one arg: url")
	}

	url := cmd.args[0]

	params := database.DeleteFeedFollowParams{
		UserID: user.ID,
		Url:    url,
	}
	err := s.db.DeleteFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Println("The feed has been unfollowed")

	return nil
}
