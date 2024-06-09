package enabler

import (
	"context"
	"fmt"
	"github.com/Abdulsametileri/go-binary-version-manager/internal/commandrunner"
	"os"
)

type golangciLintVersionEnabler struct {
	goCmdRunner commandrunner.CommandRunner
	os          OS
}

var _ VersionEnabler = (*golangciLintVersionEnabler)(nil)

func (g *golangciLintVersionEnabler) Enable(_ context.Context, version string) error {
	goRootPath, err := g.goCmdRunner.RunWith("env", "GOPATH")
	if err != nil {
		return err
	}

	// we can safely ignore, its not important path exist or not
	_ = g.os.Remove(fmt.Sprintf("%s/bin/golangci-lint", goRootPath))

	oldName := fmt.Sprintf("%s/bin/glvm/golangci-lint/%s/golangci-lint", goRootPath, version)

	if _, err = g.os.Stat(oldName); os.IsNotExist(err) {
		return fmt.Errorf("golangci-lint version %s is not exist, you can install it first", version)
	}

	newName := fmt.Sprintf("%s/bin/golangci-lint", goRootPath)

	return g.os.Symlink(oldName, newName)
}
