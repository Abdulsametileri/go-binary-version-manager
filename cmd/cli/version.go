package cli

import (
	"encoding/json"
	"fmt"
	"runtime"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	GitCommitSHA        = "unknown"
	BuildDate           = "unknown"
	Version      string = "0.0.1"
)

type VersionInfo struct {
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
		RunE: func(_ *cobra.Command, _ []string) error {
			vi := &VersionInfo{
				Version:      Version,
				GitCommitSHA: GitCommitSHA,
				BuildDate:    BuildDate,
				GoVersion:    runtime.Version(),
				Compiler:     runtime.Compiler,
				Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
			}
			bytes, err := json.Marshal(vi)
			if err != nil {
				return fmt.Errorf("failed to marshal version info %s", err.Error())
			}

			log.Infof(string(bytes))
			return nil
		},
	}
	return cmd
}
