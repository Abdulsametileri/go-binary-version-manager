package main

import (
	"github.com/Abdulsametileri/go-binary-version-manager/cmd"
	"os"
)

func main() {
	if err := cmd.New().Execute(); err != nil {
		os.Exit(1)
	}
}
