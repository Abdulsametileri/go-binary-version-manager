package cmd

import (
	"io"
	"os"

	"github.com/Abdulsametileri/go-binary-version-manager/cmd/cli"
	"github.com/Abdulsametileri/go-binary-version-manager/cmd/cli/options"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ro = &options.RootOptions{}

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gbvm",
		Short: "Manage version of libraries which installed via `go install` command.",
		PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
			return setUpLogs(os.Stdout, log.Level(ro.Verbosity))
		},
		Run: func(cmd *cobra.Command, _ []string) {
			_ = cmd.Help()
		},
	}

	ro.AddFlags(cmd)

	cmd.AddCommand(cli.InstallCmd())
	cmd.AddCommand(cli.EnableCmd())
	cmd.AddCommand(cli.ListAllCmd())
	cmd.AddCommand(cli.UnInstallCmd())
	cmd.AddCommand(cli.VersionCmd())

	return cmd
}

// setUpLogs set the log output ans the log level
func setUpLogs(out io.Writer, level log.Level) error {
	log.SetOutput(out)
	log.SetOutput(os.Stdout)
	lvl, err := log.ParseLevel(level.String())
	if err != nil {
		return err
	}
	log.SetLevel(lvl)
	return nil
}
