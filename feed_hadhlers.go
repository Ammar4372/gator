package main

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/Ammar4372/gator/internal/database"
	"github.com/google/uuid"
)

func addFeedHandler(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %v <name> <url>", cmd.name)
	}
	user, err := s.db.GetUser(context.Background(), s.cfg.UserName)
	if err != nil {
		return fmt.Errorf("couldn't find user %w", err)
	}
	_, err = url.Parse(cmd.args[1])
	if err != nil {
		return fmt.Errorf("error parsing url %s: %w", cmd.args[1], err)
	}
	params := database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	feed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}
	fmt.Printf("Name: %s\nUrl: %s", feed.Name, feed.Url)
	return nil
}
func listFeedsHandler(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get feeds: %w", err)
	}
	for _, v := range feeds {
		user, err := s.db.GetUserByID(context.Background(), v.UserID)
		if err != nil {
			return fmt.Errorf("couldn't get user: %w", err)
		}
		fmt.Printf("Name: %s\nUrl: %s\nCreated By: %s\n", v.Name, v.CreatedAt, user.Name)
	}
	return nil
}
