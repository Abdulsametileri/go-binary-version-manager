package internal

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/Abdulsametileri/go-binary-version-manager/internal/commandrunner"
	"github.com/Abdulsametileri/go-binary-version-manager/pkg"
)

type Uninstaller interface {
	Uninstall(ctx context.Context, lib, version string) error
}

type uninstaller struct {
	goCmdRunner commandrunner.CommandRunner
	os          pkg.OS
}

func NewUninstaller(runner commandrunner.CommandRunner, os pkg.OS) Uninstaller {
	return &uninstaller{
		goCmdRunner: runner,
		os:          os,
	}
}

var _ Uninstaller = (*uninstaller)(nil)

func (u *uninstaller) Uninstall(_ context.Context, lib, version string) error {
	goRootPath, err := u.goCmdRunner.RunWith("env", "GOPATH")
	if err != nil {
		return err
	}
	path := fmt.Sprintf("%s/bin/glvm/%s/%s", goRootPath, lib, version)
	if err = u.os.RemoveAll(path); err != nil {
		return err
	}

	log.Infof("%s@%s's path %s is removed successfully", lib, version, path)
	return nil
}
