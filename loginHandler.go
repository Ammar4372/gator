package main

import (
	"context"
	"fmt"
)

func loginHandler(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("invalid arguments. a username is required")
	}
	ctx := context.Background()
	_, err := s.db.GetUser(ctx, cmd.args[0])
	if err != nil {
		return fmt.Errorf("User Not found")
	}
	s.cfg.SetUser(cmd.args[0])
	fmt.Printf("Successfully loggedin as %s\n", cmd.args[0])
	return nil
}
