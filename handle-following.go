package main

import (
	"context"
	"errors"
	"fmt"
	"log"
)

func handleFollowing(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		log.Fatal(errors.New("This command doesn't take any args"))
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		log.Fatal(err)
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("You are following these feeds:")

	for _, feedFollow := range feedFollows {
		fmt.Printf("name: %s, user: %s", feedFollow.FeedName, feedFollow.UserName)
	}

	return nil
}
