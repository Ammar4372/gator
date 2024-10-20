package main

import (
	"fmt"
	"os"

	"github.com/Ammar4372/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
	err = cfg.SetUser("Ammar")
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
	cfg, err = config.Read()
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, cfg.DBUrl)
	fmt.Fprintln(os.Stdout, cfg.UserName)
}
