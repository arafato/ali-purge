package main

import (
	"os"

	"github.com/arafato/ali-purge/cmd"
)

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
