package cli

import (
	"context"

	"github.com/Abdulsametileri/go-binary-version-manager/cmd/cli/options"
	"github.com/Abdulsametileri/go-binary-version-manager/internal"
	"github.com/Abdulsametileri/go-binary-version-manager/internal/commandrunner"
	"github.com/Abdulsametileri/go-binary-version-manager/pkg"
	"github.com/spf13/cobra"
)

func UnInstallCmd() *cobra.Command {
	o := &options.UnInstallOptions{}

	cmd := &cobra.Command{
		Use:          "uninstall",
		Short:        "it uninstalls given version of the library",
		SilenceUsage: true,
		PreRunE: func(_ *cobra.Command, args []string) error {
			return o.SetLibraryNameAndVersion(args)
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return UnInstall(cmd.Context(), o)
		},
	}

	return cmd
}

func UnInstall(ctx context.Context, o *options.UnInstallOptions) error {
	uninstaller := internal.NewUninstaller(commandrunner.Get("go"), pkg.RealOs{})
	return uninstaller.Uninstall(ctx, o.LibraryName, o.Version)
}
