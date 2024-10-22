package main

import (
	"context"
	"fmt"
)

func aggHandler(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couldn't fetch feed %w", err)
	}
	fmt.Printf("Feed Title %s \n", feed.Channel.Title)
	fmt.Printf("Feed Link %s \n", feed.Channel.Link)
	fmt.Printf("Feed Description %s \n", feed.Channel.Description)
	for k, v := range feed.Channel.Item {
		fmt.Printf("Item %d Link %s\n", k, v.Link)
		fmt.Printf("Item %d Title %s\n", k, v.Title)
		fmt.Printf("Item %d Description %s\n", k, v.Description)
		fmt.Printf("Item %d Publish Date %s\n", k, v.PubDate)
	}
	return nil
}
