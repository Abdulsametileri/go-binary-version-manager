package cli

import (
	"context"
	"github.com/Abdulsametileri/go-binary-version-manager/cmd/cli/options"
	"github.com/Abdulsametileri/go-binary-version-manager/internal/installer"
	"github.com/spf13/cobra"
)

func InstallCmd() *cobra.Command {
	o := &options.InstallOptions{}

	cmd := &cobra.Command{
		Use:          "install",
		Short:        "it installs given version of the library",
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return o.SetLibraryNameAndVersion(args)
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return Install(cmd.Context(), o)
		},
	}

	return cmd
}

func Install(ctx context.Context, o *options.InstallOptions) error {
	inst, err := installer.Get(o.LibraryName)
	if err != nil {
		return err
	}

	return inst.Install(ctx, o.Version)
}
