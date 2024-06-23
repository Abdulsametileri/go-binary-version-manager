package cli

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/Abdulsametileri/go-binary-version-manager/cmd/cli/options"
	"github.com/Abdulsametileri/go-binary-version-manager/pkg"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func ListAllCmd() *cobra.Command {
	o := &options.ListAllOptions{}

	cmd := &cobra.Command{
		Use:          "listall",
		Short:        "it lists all installed versions of the given library",
		SilenceUsage: true,
		PreRunE: func(_ *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("provide a library name")
			}
			o.LibName = args[0]
			return nil
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return ListAll(cmd.Context(), o)
		},
	}

	return cmd
}

func ListAll(_ context.Context, o *options.ListAllOptions) error {
	goRootPath, err := pkg.RunCommand("go", "env", "GOPATH")
	if err != nil {
		return err
	}

	log.Debugf("go env GOPATH output %s", goRootPath)

	root := fmt.Sprintf("%s/bin/glvm/%s", goRootPath, o.LibName)

	log.Debugf("Searching path %s", root)
	buf := bytes.Buffer{}
	err = filepath.Walk(root, func(_ string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && info.Name() != o.LibName && info.Name() != "glvm" {
			buf.WriteString(info.Name())
			buf.WriteString("\n")
		}
		return nil
	})
	if err != nil {
		return err
	}

	log.Infof("Getting all files within %s", root)
	_, err = buf.WriteTo(os.Stdout)
	return err
}
