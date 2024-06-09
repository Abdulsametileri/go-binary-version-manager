package commandrunner

import (
	log "github.com/sirupsen/logrus"
)

//go:generate mockery --name=CommandRunner --filename=command_runner.go --output=../installer/mocks
type CommandRunner interface {
	Run(cmd string) ([]byte, error)
}

func Get(cmd string) CommandRunner {
	if cmd == "curl" {
		return &curlCommandRunner{}
	}
	log.Fatalf("command runner %s not found", cmd)
	return nil
}
