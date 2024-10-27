package main

import (
	"context"
	"errors"
	"fmt"
)

func handleFeeds(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("This command doesn't accept any args")
	}

	_, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feeds, err := s.db.GetFeedsWithUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Here are all of the feeds:")
	fmt.Println()
	for _, feed := range feeds {
		fmt.Printf("name: %s, url: %s, user: %s\n", feed.Name, feed.Url, feed.Name_2)
	}

	return nil
}
