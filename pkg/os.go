package pkg

import "os"

//go:generate mockery --name=OS --filename=os.go --output=../internal/mocks
type OS interface {
	Remove(name string) error
	RemoveAll(path string) error
	Stat(name string) (os.FileInfo, error)
	Symlink(oldname, newname string) error
}

var _ OS = (*RealOs)(nil)

type RealOs struct{}

func (RealOs) RemoveAll(path string) error {
	return os.RemoveAll(path)
}

func (RealOs) Remove(name string) error {
	return os.Remove(name)
}

func (RealOs) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

func (RealOs) Symlink(oldname, newname string) error {
	return os.Symlink(oldname, newname)
}
