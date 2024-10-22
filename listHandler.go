package main

import (
	"context"
	"fmt"
)

func listHandler(s *state, cmd command) error {

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get user %w", err)
	}
	for _, user := range users {
		if user.Name == s.cfg.UserName {
			fmt.Printf("%s (current)\n", user.Name)
			continue
		}
		fmt.Printf("%s\n", user.Name)
	}
	return nil
}
