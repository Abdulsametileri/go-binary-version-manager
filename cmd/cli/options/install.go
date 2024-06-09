package options

import (
	"errors"
	"fmt"
	"strings"
)

type InstallOptions struct {
	Library string
	Version string
}

func (io *InstallOptions) SetLibraryNameAndVersion(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("provide a library@version. Like `gobvm install mockery@v2.20.0`")
	}

	splitted := strings.Split(args[0], "@")
	if len(splitted) != 2 {
		return errors.New("provide a library@version")
	}

	io.Library, io.Version = splitted[0], splitted[1]
	return nil
}
