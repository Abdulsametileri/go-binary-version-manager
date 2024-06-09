package cli

import "github.com/spf13/cobra"

func InstallCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "it installs given version of the library",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}
