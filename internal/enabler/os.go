package enabler

import "os"

//go:generate mockery --name=OS --filename=os.go --output=./mocks
type OS interface {
	Remove(name string) error
	Stat(name string) (os.FileInfo, error)
	Symlink(oldname, newname string) error
}

var _ OS = (*realOs)(nil)

type realOs struct{}

func (realOs) Remove(name string) error {
	return os.Remove(name)
}

func (realOs) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

func (realOs) Symlink(oldname, newname string) error {
	return os.Symlink(oldname, newname)
}
