package main

import (
	"fmt"
	"stash.notes/internal/config"
	"stash.notes/internal/ui/tui"
)

func main() {
	fmt.Println("Hello, from Stash.Notes")

	config.LoadConfig()

	tui.RunTui()
}
