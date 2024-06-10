package internal

import (
	"context"
	"errors"
	"github.com/Abdulsametileri/go-binary-version-manager/internal/mocks"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_golangciLintVersionEnabler_Enable(t *testing.T) {
	t.Run("Should_Return_Error_When_Go_Command_Failed", func(t *testing.T) {
		// Given
		expectedErr := errors.New("command err")
		runner := mocks.NewCommandRunner(t)
		runner.On("RunWith", "env", "GOPATH").Return("", expectedErr)
		sut := &versionEnabler{goCmdRunner: runner}

		// When
		result := sut.Enable(context.Background(), "golangci-lint", "v1.55.0")

		// Then
		assert.ErrorIs(t, result, expectedErr)
		runner.AssertExpectations(t)
	})
	t.Run("Should_Return_Error_Specified_Version_Does_Not_Exist", func(t *testing.T) {
		// Given
		runner := mocks.NewCommandRunner(t)
		runner.On("RunWith", "env", "GOPATH").Return("/Users/samet.ileri/go", nil)

		mockOs := mocks.NewOS(t)
		mockOs.On("Stat", "/Users/samet.ileri/go/bin/glvm/golangci-lint/v1.55.0/golangci-lint").Return(nil, os.ErrNotExist)

		sut := &versionEnabler{goCmdRunner: runner, os: mockOs}

		// When
		result := sut.Enable(context.Background(), "golangci-lint", "v1.55.0")

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
		mockOs.On("Stat", "/Users/samet.ileri/go/bin/glvm/golangci-lint/v1.55.0/golangci-lint").Return(MockFileInfo{}, nil)
		mockOs.On("Remove", "/Users/samet.ileri/go/bin/golangci-lint").Return(nil)
		mockOs.
			On("Symlink",
				"/Users/samet.ileri/go/bin/glvm/golangci-lint/v1.55.0/golangci-lint",
				"/Users/samet.ileri/go/bin/golangci-lint").
			Return(nil)

		sut := &versionEnabler{goCmdRunner: runner, os: mockOs}

		// When
		result := sut.Enable(context.Background(), "golangci-lint", "v1.55.0")

		// Then
		assert.NoError(t, result)
		runner.AssertExpectations(t)
		mockOs.AssertExpectations(t)
	})
}
