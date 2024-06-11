package installer

import (
	"context"
	"fmt"
	"github.com/Abdulsametileri/go-binary-version-manager/internal/commandrunner"
	log "github.com/sirupsen/logrus"
	"os"
)

type mockeryInstaller struct {
	goCmdRunner commandrunner.CommandRunner
	goBinEnvKey string
}

func (m *mockeryInstaller) Install(_ context.Context, version string) error {
	goRootPath, err := m.goCmdRunner.RunWith("env", "GOPATH")
	if err != nil {
		return err
	}

	goBinOriginal := os.Getenv(m.goBinEnvKey)
	defer os.Setenv(m.goBinEnvKey, goBinOriginal)

	if err = os.Setenv(m.goBinEnvKey, fmt.Sprintf("%s/bin/glvm/mockery/%s", goRootPath, version)); err != nil {
		return fmt.Errorf("error setting %s environment variable %w", m.goBinEnvKey, err)
	}

	out, err := m.goCmdRunner.RunWith("install", "github.com/vektra/mockery/v2@"+version)
	if err != nil {
		return err
	}
	log.Infof("Command Output: %s\n", out)
	return nil
}
