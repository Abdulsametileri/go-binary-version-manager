package cli

import (
	"context"
	"errors"
	"fmt"
	"github.com/Abdulsametileri/go-binary-version-manager/cmd/cli/options"
	"github.com/Abdulsametileri/go-binary-version-manager/internal/lister"
	"github.com/Abdulsametileri/go-binary-version-manager/internal/model"
	"github.com/spf13/cobra"
)

func ListAllCmd() *cobra.Command {
	o := &options.ListAllOptions{}

	cmd := &cobra.Command{
		Use:          "listall",
		Short:        "it lists all installed versions of the given library",
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("provide a library name")
			}
			o.Library = args[0]

			if _, supported := model.SupportedLibraries[model.Library(o.Library)]; !supported {
				return fmt.Errorf("lib %s is not supported", o.Library)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return ListAll(cmd.Context(), o)
		},
	}

	return cmd
}

func ListAll(ctx context.Context, o *options.ListAllOptions) error {
	stdoutLister, err := lister.Get("stdout")
	if err != nil {
		return err
	}

	return stdoutLister.List(ctx, o.Library)
}
