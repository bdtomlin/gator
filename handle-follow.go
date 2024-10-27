package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/bdtomlin/gator/internal/database"
	"github.com/google/uuid"
)

func handleFollow(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		log.Fatal(errors.New("This command requires one arg: url"))
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		log.Fatal(err)
	}

	url := cmd.args[0]
	feed, err := s.db.GetFeedForUrl(context.Background(), url)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	fmt.Println("The feed has been followed")
	fmt.Printf("%+v", feedFollow)

	return nil
}
