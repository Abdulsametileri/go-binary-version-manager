package cli

import (
	"context"

	"github.com/Abdulsametileri/go-binary-version-manager/cmd/cli/options"
	"github.com/Abdulsametileri/go-binary-version-manager/internal"
	"github.com/Abdulsametileri/go-binary-version-manager/internal/commandrunner"
	"github.com/Abdulsametileri/go-binary-version-manager/pkg"
	"github.com/spf13/cobra"
)

func EnableCmd() *cobra.Command {
	o := &options.EnableOptions{}

	cmd := &cobra.Command{
		Use:          "enable",
		Short:        "it enables given version of the library",
		SilenceUsage: true,
		PreRunE: func(_ *cobra.Command, args []string) error {
			return o.SetLibraryNameAndVersion(args)
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return Enable(cmd.Context(), o)
		},
	}

	return cmd
}

func Enable(ctx context.Context, o *options.EnableOptions) error {
	enabler := internal.NewVersionEnabler(commandrunner.Get("go"), pkg.RealOs{})
	return enabler.Enable(ctx, o.LibraryName, o.Version)
}
