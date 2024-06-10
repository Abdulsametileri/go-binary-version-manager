package lister

import (
	"context"
	"fmt"
	"github.com/Abdulsametileri/go-binary-version-manager/internal/commandrunner"
	"os"
)

type Lister interface {
	List(ctx context.Context, lib string) error
}

func Get(t string) (Lister, error) {
	switch t {
	case "stdout":
		return &stdoutLister{
			goCmdRunner: commandrunner.Get("go"), listTo: os.Stdout, walker: realFileWalker{}}, nil
	default:
		return nil, fmt.Errorf("lister %s is not supported", t)
	}
}
