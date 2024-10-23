package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Ammar4372/gator/internal/database"
	"github.com/google/uuid"
)

func registerHandler(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.name)
	}
	ctx := context.Background()

	params := database.CreateUserParams{ID: uuid.New(), Name: cmd.args[0], CreatedAt: time.Now(), UpdatedAt: time.Now()}
	user, err := s.db.CreateUser(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set user: %w", err)
	}
	fmt.Printf("User Created Successfully\n")
	return nil
}
