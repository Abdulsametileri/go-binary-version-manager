package enabler

import (
	"context"
	"errors"
	"github.com/Abdulsametileri/go-binary-version-manager/internal/enabler/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
	"time"
)

func Test_golangciLintVersionEnabler_Enable(t *testing.T) {
	t.Run("Should_Return_Error_When_Go_Command_Failed", func(t *testing.T) {
		// Given
		expectedErr := errors.New("command err")
		runner := mocks.NewCommandRunner(t)
		runner.On("RunWith", "env", "GOPATH").Return("", expectedErr)
		sut := &golangciLintVersionEnabler{goCmdRunner: runner}

		// When
		result := sut.Enable(context.Background(), "v1.55.0")

		// Then
		assert.ErrorIs(t, result, expectedErr)
		runner.AssertExpectations(t)
	})
	t.Run("Should_Return_Error_Specified_Version_Does_Not_Exist", func(t *testing.T) {
		// Given
		runner := mocks.NewCommandRunner(t)
		runner.On("RunWith", "env", "GOPATH").Return("/Users/samet.ileri/go", nil)

		mockOs := mocks.NewOS(t)
		mockOs.On("Remove", "/Users/samet.ileri/go/bin/golangci-lint").Return(nil)
		mockOs.On("Stat", "/Users/samet.ileri/go/bin/glvm/golangci-lint/v1.55.0/golangci-lint").Return(nil, os.ErrNotExist)

		sut := &golangciLintVersionEnabler{goCmdRunner: runner, os: mockOs}

		// When
		result := sut.Enable(context.Background(), "v1.55.0")

		// Then
		assert.Equal(t, result.Error(), "golangci-lint version v1.55.0 is not exist, you can install it first")
		runner.AssertExpectations(t)
		mockOs.AssertExpectations(t)
	})
	t.Run("Success", func(t *testing.T) {
		// Given
		runner := mocks.NewCommandRunner(t)
		runner.On("RunWith", "env", "GOPATH").Return("/Users/samet.ileri/go", nil)

		mockOs := mocks.NewOS(t)
		mockOs.On("Remove", "/Users/samet.ileri/go/bin/golangci-lint").Return(nil)
		mockOs.On("Stat", "/Users/samet.ileri/go/bin/glvm/golangci-lint/v1.55.0/golangci-lint").Return(MockFileInfo{}, nil)
		mockOs.
			On("Symlink",
				"/Users/samet.ileri/go/bin/glvm/golangci-lint/v1.55.0/golangci-lint",
				"/Users/samet.ileri/go/bin/golangci-lint").
			Return(nil)

		sut := &golangciLintVersionEnabler{goCmdRunner: runner, os: mockOs}

		// When
		result := sut.Enable(context.Background(), "v1.55.0")

		// Then
		assert.NoError(t, result)
		runner.AssertExpectations(t)
		mockOs.AssertExpectations(t)
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
