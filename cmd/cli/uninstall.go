package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/Abdulsametileri/go-binary-version-manager/cmd/cli/options"
	"github.com/Abdulsametileri/go-binary-version-manager/pkg"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func UnInstallCmd() *cobra.Command {
	o := &options.UnInstallOptions{}

	cmd := &cobra.Command{
		Use:          "uninstall",
		Short:        "it uninstalls given version of the library",
		SilenceUsage: true,
		PreRunE: func(_ *cobra.Command, args []string) error {
			return o.Set(args)
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return UnInstall(cmd.Context(), o)
		},
	}

	return cmd
}

func UnInstall(_ context.Context, o *options.UnInstallOptions) error {
	goRootPath, err := pkg.RunCommand("go", "env", "GOPATH")
	if err != nil {
		return err
	}

	log.Debugf("go env GOPATH output %s", goRootPath)

	path := fmt.Sprintf("%s/bin/glvm/%s/%s", goRootPath, o.LibName, o.Version)
	log.Debugf("will remove this %s", path)

	if err = os.RemoveAll(path); err != nil {
		return err
	}

	log.Infof("%s@%s's path %s is removed successfully", o.LibName, o.Version, path)
	return nil
}
