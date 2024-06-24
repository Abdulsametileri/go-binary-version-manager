package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/Abdulsametileri/go-binary-version-manager/pkg"
	log "github.com/sirupsen/logrus"

	"github.com/Abdulsametileri/go-binary-version-manager/cmd/cli/options"
	"github.com/spf13/cobra"
)

func EnableCmd() *cobra.Command {
	o := &options.EnableOptions{}

	cmd := &cobra.Command{
		Use:          "enable",
		Short:        "it enables given version of the library",
		SilenceUsage: true,
		PreRunE: func(_ *cobra.Command, args []string) error {
			return o.Set(args)
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return Enable(cmd.Context(), o)
		},
	}

	return cmd
}

func Enable(_ context.Context, o *options.EnableOptions) error {
	goRootPath, err := pkg.RunCommand("go", "env", "GOPATH")
	if err != nil {
		return err
	}

	log.Debugf("go env GOPATH output %s", goRootPath)

	oldName := fmt.Sprintf("%s/bin/glvm/%s/%s/%s", goRootPath, o.LibName, o.Version, o.LibName)

	if _, err = os.Stat(oldName); os.IsNotExist(err) {
		return fmt.Errorf("%s version %s is not exist, you can install it first", o.LibName, o.Version)
	}

	log.Debugf("%s@%s found on %s", o.LibName, o.Version, oldName)

	// we can safely ignore, its not important path exist or not
	_ = os.Remove(fmt.Sprintf("%s/bin/%s", goRootPath, o.LibName))

	newName := fmt.Sprintf("%s/bin/%s", goRootPath, o.LibName)

	if err = os.Symlink(oldName, newName); err != nil {
		return err
	}

	log.Infof("Symlink successfully added for path %s", newName)
	return nil
}
