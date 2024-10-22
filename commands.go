package main

import (
	"fmt"
	"os"
)

type commands struct {
	handlers map[string]func(state *state, cmd command) error
}

func (c *commands) register(name string, h func(state *state, cmd command) error) {
	c.handlers[name] = h
}

func (c *commands) run(state *state, cmd command) error {
	if _, ok := c.handlers[cmd.name]; ok {
		err := c.handlers[cmd.name](state, cmd)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		fmt.Println("invalid command")
		os.Exit(1)
	}
	return nil
}
