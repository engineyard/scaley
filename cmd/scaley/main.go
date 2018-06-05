package main

import (
	"os"

	"github.com/engineyard/scaley/cmd/scaley/cmd"
)

func main() {
	if cmd.Execute() != nil {
		os.Exit(1)
	}
}
