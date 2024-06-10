package pkg

import "path/filepath"

//go:generate mockery --name=FileWalker --filename=file_walker.go --output=../internal/mocks
type FileWalker interface {
	Walk(root string, fn filepath.WalkFunc) error
}

var _ FileWalker = (*RealFileWalker)(nil)

type RealFileWalker struct{}

func (RealFileWalker) Walk(root string, fn filepath.WalkFunc) error {
	return filepath.Walk(root, fn)
}
