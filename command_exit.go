package main

import (
	"fmt"
	"os"
)

func callbackExit(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Bye Snoot. I love you! 😘")
	fmt.Println()
	os.Exit(0)
	return nil
}