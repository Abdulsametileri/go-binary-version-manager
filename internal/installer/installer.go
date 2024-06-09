package installer

import (
	"context"
	"fmt"
	"github.com/Abdulsametileri/go-binary-version-manager/internal/commandrunner"
	"github.com/Abdulsametileri/go-binary-version-manager/internal/model"
)

type Installer interface {
	Install(ctx context.Context, version string) error
}

func Get(lib string) (Installer, error) {
	switch model.Library(lib) {
	case model.LibraryGolangciLint:
		return &golangciLintInstaller{commandRunner: commandrunner.Get("curl")}, nil
	default:
		return nil, fmt.Errorf("lib %s is not supported", lib)
	}
}
