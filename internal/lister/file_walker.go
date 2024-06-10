package lister

import "path/filepath"

//go:generate mockery --name=FileWalker --filename=file_walker.go --output=./mocks
type FileWalker interface {
	Walk(root string, fn filepath.WalkFunc) error
}

var _ FileWalker = (*realFileWalker)(nil)

type realFileWalker struct{}

func (realFileWalker) Walk(root string, fn filepath.WalkFunc) error {
	return filepath.Walk(root, fn)
}
