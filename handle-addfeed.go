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

func handleAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		log.Fatal(errors.New("This command requires the following args: name, url"))
	}

	name := cmd.args[0]
	url := cmd.args[1]

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		log.Fatal(err)
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	_, err = s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The feed has been added")
	fmt.Printf("%+v", feed)

	return nil
}
