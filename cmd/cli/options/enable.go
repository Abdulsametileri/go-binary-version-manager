package options

import (
	"errors"
	"fmt"
	"strings"
)

type EnableOptions struct {
	LibName string
	Version string
}

func (eo *EnableOptions) Set(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("provide a library@version. Like `gbvm enable mockery@v2.20.0`")
	}

	splitted := strings.Split(args[0], "@")
	if len(splitted) != 2 {
		return errors.New("provide a library@version")
	}

	eo.LibName, eo.Version = splitted[0], splitted[1]
	return nil
}
