package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Ammar4372/gator/internal/database"
	"github.com/google/uuid"
)

func registerHandler(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("no username is given")
	}
	ctx := context.Background()
	_, err := s.db.GetUser(ctx, cmd.args[0])
	if err == nil {
		return fmt.Errorf("User %s Already Exists", cmd.args[0])
	}
	params := database.CreateUserParams{ID: uuid.New(), Name: cmd.args[0], CreatedAt: time.Now(), UpdatedAt: time.Now()}
	user, err := s.db.CreateUser(ctx, params)
	if err != nil {
		return err
	}
	s.cfg.SetUser(user.Name)
	fmt.Printf("Register User %s Successfully\n", user.Name)
	return nil
}
