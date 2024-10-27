package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/bdtomlin/gator/internal/database"
	"github.com/google/uuid"
)

func handleFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return errors.New("This command requires one arg: url")
	}

	url := cmd.args[0]
	feed, err := s.db.GetFeedForUrl(context.Background(), url)
	if err != nil {
		return err
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return err
	}

	fmt.Println("The feed has been followed")
	fmt.Printf("%+v", feedFollow)

	return nil
}
