package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/bdtomlin/gator/internal/database"
)

func handleFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 0 {
		return errors.New("This command doesn't take any args")
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	fmt.Println("You are following these feeds:")

	for _, feedFollow := range feedFollows {
		fmt.Printf("name: %s, user: %s", feedFollow.FeedName, feedFollow.UserName)
	}

	return nil
}
