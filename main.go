package main

import (
	"os"

	"github.com/Abdulsametileri/go-binary-version-manager/cmd"
)

func main() {
	if err := cmd.New().Execute(); err != nil {
		os.Exit(1)
	}
}
