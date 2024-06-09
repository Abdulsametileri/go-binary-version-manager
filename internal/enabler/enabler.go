package enabler

import (
	"context"
	"fmt"
	"github.com/Abdulsametileri/go-binary-version-manager/internal/commandrunner"
	"github.com/Abdulsametileri/go-binary-version-manager/internal/model"
)

type VersionEnabler interface {
	Enable(ctx context.Context, version string) error
}

func Get(lib string) (VersionEnabler, error) {
	switch model.Library(lib) {
	case model.LibraryGolangciLint:
		return &golangciLintVersionEnabler{
			goCmdRunner: commandrunner.Get("go"),
			os:          &realOs{},
		}, nil
	default:
		return nil, fmt.Errorf("lib %s is not supported for version enabling", lib)
	}
}
