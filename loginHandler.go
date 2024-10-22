package main

import (
	"context"
	"fmt"
)

func loginHandler(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %v <name>", cmd.name)
	}
	ctx := context.Background()
	_, err := s.db.GetUser(ctx, cmd.args[0])
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}
	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("couldn't set user: %w", err)
	}
	fmt.Printf("Successfully loggedin as %s\n", cmd.args[0])
	return nil
}
