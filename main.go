package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Ammar4372/gator/internal/config"
	"github.com/Ammar4372/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

type command struct {
	name string
	args []string
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	dbQueries := database.New(db)
	s := state{
		cfg: &cfg,
		db:  dbQueries,
	}
	cmds := commands{
		handlers: map[string]func(state *state, cmd command) error{},
	}
	cmds.register("login", loginHandler)
	cmds.register("register", registerHandler)
	cmds.register("reset", resetHandler)
	cmds.register("users", listHandler)
	cmds.register("agg", aggHandler)
	cmds.register("addfeed", addFeedHandler)
	cmds.register("feeds", listFeedsHandler)
	args := os.Args
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "no command name given")
		os.Exit(1)
	}
	cmd := command{
		name: args[1],
		args: args[2:],
	}
	err = cmds.run(&s, cmd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
