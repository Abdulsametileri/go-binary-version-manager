package internal

import (
	"context"
	"fmt"
	"github.com/Abdulsametileri/go-binary-version-manager/internal/commandrunner"
	"github.com/Abdulsametileri/go-binary-version-manager/pkg"
	"os"
)

type VersionEnabler interface {
	Enable(ctx context.Context, lib, version string) error
}

type versionEnabler struct {
	goCmdRunner commandrunner.CommandRunner
	os          pkg.OS
}

var _ VersionEnabler = (*versionEnabler)(nil)

func NewVersionEnabler(runner commandrunner.CommandRunner, os pkg.OS) VersionEnabler {
	return &versionEnabler{
		goCmdRunner: runner,
		os:          os,
	}
}

func (g *versionEnabler) Enable(_ context.Context, lib, version string) error {
	goRootPath, err := g.goCmdRunner.RunWith("env", "GOPATH")
	if err != nil {
		return err
	}

	oldName := fmt.Sprintf("%s/bin/glvm/%s/%s/%s", goRootPath, lib, version, lib)

	if _, err = g.os.Stat(oldName); os.IsNotExist(err) {
		return fmt.Errorf("%s version %s is not exist, you can install it first", lib, version)
	}

	// we can safely ignore, its not important path exist or not
	_ = g.os.Remove(fmt.Sprintf("%s/bin/%s", goRootPath, lib))

	newName := fmt.Sprintf("%s/bin/%s", goRootPath, lib)

	return g.os.Symlink(oldName, newName)
}
