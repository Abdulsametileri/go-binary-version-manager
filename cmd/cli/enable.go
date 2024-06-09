package cli

import (
	"context"
	"github.com/Abdulsametileri/go-binary-version-manager/cmd/cli/options"
	"github.com/Abdulsametileri/go-binary-version-manager/internal/enabler"
	"github.com/spf13/cobra"
)

func EnableCmd() *cobra.Command {
	o := &options.EnableOptions{}

	cmd := &cobra.Command{
		Use:          "enable",
		Short:        "it enables given version of the library",
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return o.SetLibraryNameAndVersion(args)
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return Enable(cmd.Context(), o)
		},
	}

	return cmd
}

func Enable(ctx context.Context, o *options.EnableOptions) error {
	enablr, err := enabler.Get(o.LibraryName)
	if err != nil {
		return err
	}

	return enablr.Enable(ctx, o.Version)
}
