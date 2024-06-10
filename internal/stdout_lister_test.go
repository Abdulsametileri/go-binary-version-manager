package internal

import (
	"bytes"
	"context"
	"errors"
	"github.com/Abdulsametileri/go-binary-version-manager/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func Test_stdoutLister_List(t *testing.T) {
	t.Run("Should_Return_Error_When_Go_Command_Failed", func(t *testing.T) {
		// Given
		expectedErr := errors.New("command err")
		runner := mocks.NewCommandRunner(t)
		runner.On("RunWith", "env", "GOPATH").Return("", expectedErr)

		sut := &stdoutLister{
			goCmdRunner: runner,
		}

		// When
		err := sut.List(context.Background(), "mockery")

		// Then
		assert.ErrorIs(t, err, expectedErr)
	})
	t.Run("Should_Return_Error_When_FileWalker_Failed", func(t *testing.T) {
		// Given
		expectedErr := errors.New("walker err")
		runner := mocks.NewCommandRunner(t)
		mockWalker := mocks.NewFileWalker(t)

		runner.On("RunWith", "env", "GOPATH").Return("/Users/samet.ileri/go", nil)
		mockWalker.On("Walk", mock.Anything, mock.Anything).Return(expectedErr)

		sut := &stdoutLister{
			goCmdRunner: runner,
			walker:      mockWalker,
		}

		// When
		err := sut.List(context.Background(), "mockery")

		// Then
		assert.ErrorIs(t, err, expectedErr)
	})
	t.Run("Success", func(t *testing.T) {
		// Given
		runner := mocks.NewCommandRunner(t)
		mockWalker := mocks.NewFileWalker(t)

		runner.On("RunWith", "env", "GOPATH").Return("/Users/samet.ileri/go", nil)
		mockWalker.
			On("Walk", "/Users/samet.ileri/go/bin/glvm/mockery", mock.Anything).
			Return(nil).
			Run(func(args mock.Arguments) {
				fn := args.Get(1).(filepath.WalkFunc)
				fn("/Users/samet.ileri/go/bin/glvm/mockery/v2.20.0", MockFileInfo{name: "v2.20.0", isDir: true}, nil)
				fn("/Users/samet.ileri/go/bin/glvm/mockery/v2.25.0", MockFileInfo{name: "v2.25.0", isDir: true}, nil)
				fn("/Users/samet.ileri/go/bin/glvm/mockery/v2.25.0", MockFileInfo{name: "mockery", isDir: false}, nil)
				fn("/Users/samet.ileri/go/bin/glvm/mockery/v2.29.0", MockFileInfo{name: "v2.29.0", isDir: true}, errors.New("some err"))
			})

		outputBuffer := new(bytes.Buffer)

		sut := &stdoutLister{
			goCmdRunner: runner,
			walker:      mockWalker,
			listTo:      outputBuffer,
		}

		// When
		err := sut.List(context.Background(), "mockery")

		// Then
		assert.Contains(t, outputBuffer.String(), "v2.20.0")
		assert.Contains(t, outputBuffer.String(), "v2.25.0")
		assert.NotContains(t, outputBuffer.String(), "v2.29.0")
		assert.Nil(t, err)
	})
}

type MockFileInfo struct {
	mock.Mock
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
	isDir   bool
	sys     interface{}
}

func (m MockFileInfo) Name() string       { return m.name }
func (m MockFileInfo) Size() int64        { return m.size }
func (m MockFileInfo) Mode() os.FileMode  { return m.mode }
func (m MockFileInfo) ModTime() time.Time { return m.modTime }
func (m MockFileInfo) IsDir() bool        { return m.isDir }
func (m MockFileInfo) Sys() interface{}   { return m.sys }
