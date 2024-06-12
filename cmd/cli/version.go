package cli

import (
	"encoding/json"
	"fmt"
	"github.com/Abdulsametileri/go-binary-version-manager/internal"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"runtime"
)

var (
	GitCommitSHA = "unknown"
	BuildDate    = "unknown"
)

type CLIVersionInfo struct {
	Version      string
	GitCommitSHA string
	BuildDate    string
	GoVersion    string
	Compiler     string
	Platform     string
}

func VersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "version",
		Short:        "Prints the CLI version",
		Long:         `Prints the CLI version`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			bytes, err := json.Marshal(VersionInfo())
			if err != nil {
				return fmt.Errorf("failed to marshal version info %s", err.Error())
			}

			log.Infof(string(bytes))
			return nil
		},
	}
	return cmd
}

func VersionInfo() *CLIVersionInfo {
	return &CLIVersionInfo{
		Version:      internal.Version,
		GitCommitSHA: GitCommitSHA,
		BuildDate:    BuildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
