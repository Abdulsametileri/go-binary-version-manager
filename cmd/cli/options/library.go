package options

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Abdulsametileri/go-binary-version-manager/internal"
)

/*
LibraryOptions for example; github.com/vektra/mockery/v2@v2.20.0
Address => github.com/vektra/mockery/v2
Version => v2.20.0
Package => github.com/vektra/mockery/v2@v2.20.0
LibName => mockery
*/
type LibraryOptions struct {
	Package string
	Address string
	Version string
	LibName string
}

func (lo *LibraryOptions) Set(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("provide a library@version. " +
			"Like `gbvm install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.0`")
	}

	libName, err := internal.ExtractLibName(args[0])
	if err != nil {
		return err
	}

	splitted := strings.Split(args[0], "@")
	if len(splitted) != 2 {
		return errors.New("provide a library-address@version")
	}

	lo.Package = args[0]
	lo.Address, lo.Version = splitted[0], splitted[1]
	lo.LibName = libName

	return nil
}
