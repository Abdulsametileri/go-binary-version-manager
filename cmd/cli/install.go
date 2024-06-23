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

func InstallCmd() *cobra.Command {
	o := &options.InstallOptions{}

	return &cobra.Command{
		Use:          "install",
		Short:        "it installs given version of the library",
		SilenceUsage: true,
		PreRunE: func(_ *cobra.Command, args []string) error {
			return o.Set(args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return Install(cmd.Context(), o)
		},
	}
}

func Install(_ context.Context, o *options.InstallOptions) error {
	goRootPath, err := pkg.RunCommand("go", "env", "GOPATH")
	if err != nil {
		return err
	}

	log.Debugf("go env GOPATH output %s", goRootPath)

	// go install command installs the binary which located in GOBIN
	// in order to configure symlink we need to change GOBIN temporary
	goBinOriginal := os.Getenv("GOBIN")
	defer os.Setenv("GOBIN", goBinOriginal)

	if err = os.Setenv("GOBIN", fmt.Sprintf("%s/bin/glvm/%s/%s", goRootPath, o.LibName, o.Version)); err != nil {
		return fmt.Errorf("error setting GOBIN environment variable %w", err)
	}

	out, err := pkg.RunCommand("go", "install", o.Package)
	if err != nil {
		return err
	}

	log.Infof("Lib %s version %s Successfully installed!\n%s", o.LibName, o.Version, out)
	return nil
}

// TODO: where to put?
