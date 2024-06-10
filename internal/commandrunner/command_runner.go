package commandrunner

import (
	log "github.com/sirupsen/logrus"
)

//go:generate mockery --name=CommandRunner --filename=command_runner.go --output=../installer/mocks
//go:generate mockery --name=CommandRunner --filename=command_runner.go --output=../enabler/mocks
//go:generate mockery --name=CommandRunner --filename=command_runner.go --output=../lister/mocks
type CommandRunner interface {
	RunWith(args ...string) (output string, err error)
}

func Get(name string) CommandRunner {
	if name == "curl" {
		return &curlCommandRunner{}
	} else if name == "go" {
		return &goCommandRunner{}
	}

	log.Fatalf("command runner %s not found", name)
	return nil
}
