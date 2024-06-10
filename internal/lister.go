package internal

import (
	"bytes"
	"context"
	"fmt"
	"github.com/Abdulsametileri/go-binary-version-manager/internal/commandrunner"
	"github.com/Abdulsametileri/go-binary-version-manager/pkg"
	"io"
	"io/fs"
)

type Lister interface {
	List(ctx context.Context, lib string) error
}

var _ Lister = (*stdoutLister)(nil)

type stdoutLister struct {
	goCmdRunner commandrunner.CommandRunner
	listTo      io.Writer
	walker      pkg.FileWalker
}

func NewStdoutLister(runner commandrunner.CommandRunner, listTo io.Writer, walker pkg.FileWalker) Lister {
	return &stdoutLister{
		goCmdRunner: runner,
		listTo:      listTo,
		walker:      walker,
	}
}

func (s *stdoutLister) List(_ context.Context, lib string) error {
	goRootPath, err := s.goCmdRunner.RunWith("env", "GOPATH")
	if err != nil {
		return err
	}

	root := fmt.Sprintf("%s/bin/glvm/%s", goRootPath, lib)

	buf := bytes.Buffer{}
	err = s.walker.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && info.Name() != lib && info.Name() != "glvm" {
			buf.WriteString(info.Name())
			buf.WriteString("\n")
		}
		return nil
	})
	if err != nil {
		return err
	}

	_, err = buf.WriteTo(s.listTo)
	return err
}
