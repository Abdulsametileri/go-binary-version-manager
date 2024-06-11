package installer

import (
	"context"
	"fmt"

	"github.com/Abdulsametileri/go-binary-version-manager/internal/commandrunner"
	log "github.com/sirupsen/logrus"
)

type golangciLintInstaller struct {
	curlCmdRunner commandrunner.CommandRunner
}

var _ Installer = (*golangciLintInstaller)(nil)

func (g *golangciLintInstaller) Install(_ context.Context, version string) error {
	curlCmd := fmt.Sprintf("curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin/glvm/golangci-lint/%s %s", version, version) //nolint:lll
	out, err := g.curlCmdRunner.RunWith(curlCmd)
	if err != nil {
		return err
	}
	log.Infof("Command Output: %s\n", out)
	return nil
}
