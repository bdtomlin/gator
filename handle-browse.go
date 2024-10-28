package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/bdtomlin/gator/internal/database"
)

func handleBrowse(s *state, cmd command, user database.User) error {
	var limit int32 = 2

	if len(cmd.args) == 1 {
		i, err := strconv.ParseInt(cmd.args[0], 10, 32)
		if err != nil {
			return errors.New("This command only accepts a single integer option")
		}
		limit = int32(i)
	}

	params := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	}
	posts, err := s.db.GetPostsForUser(context.Background(), params)
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Println("Title:", post.Title)
		fmt.Println("Url:", post.Url)
		fmt.Println("Published:", post.PublishedAt)
		fmt.Println("Description:")
		fmt.Println(post.Description)
		fmt.Println()
	}

	return nil
}
