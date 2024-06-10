package cmd

import (
	"github.com/Abdulsametileri/go-binary-version-manager/cmd/cli"
	"github.com/Abdulsametileri/go-binary-version-manager/cmd/cli/options"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var ro = &options.RootOptions{}

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gobvm",
		Short: "Manage version of go libraries. Currently `golangci-lint` and `mockery` support",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return setUpLogs(os.Stdout, log.Level(ro.Verbosity))
		},
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	ro.AddFlags(cmd)

	cmd.AddCommand(cli.InstallCmd())
	cmd.AddCommand(cli.EnableCmd())
	cmd.AddCommand(cli.ListAllCmd())
	cmd.AddCommand(cli.UnInstallCmd())

	return cmd
}

// setUpLogs set the log output ans the log level
func setUpLogs(out io.Writer, level log.Level) error {
	log.SetOutput(out)
	lvl, err := log.ParseLevel(level.String())
	if err != nil {
		return err
	}
	log.SetLevel(lvl)
	return nil
}
