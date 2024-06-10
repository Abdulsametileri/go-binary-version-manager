package internal

import (
	"context"
	"errors"
	"github.com/Abdulsametileri/go-binary-version-manager/internal/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_uninstaller_Uninstall(t *testing.T) {
	t.Run("Should_Return_Error_When_Go_Command_Failed", func(t *testing.T) {
		// Given
		expectedErr := errors.New("command err")
		runner := mocks.NewCommandRunner(t)
		runner.On("RunWith", "env", "GOPATH").Return("", expectedErr)

		sut := &uninstaller{
			goCmdRunner: runner,
		}

		// When
		err := sut.Uninstall(context.Background(), "mockery", "v2.20.0")

		// Then
		assert.ErrorIs(t, err, expectedErr)
	})
	t.Run("Success", func(t *testing.T) {
		// Given
		runner := mocks.NewCommandRunner(t)
		mockOs := mocks.NewOS(t)

		runner.On("RunWith", "env", "GOPATH").Return("/Users/samet.ileri/go", nil)
		mockOs.On("RemoveAll", "/Users/samet.ileri/go/bin/glvm/mockery/v2.20.0").Return(nil)

		sut := NewUninstaller(runner, mockOs)

		// When
		err := sut.Uninstall(context.Background(), "mockery", "v2.20.0")

		// Then
		assert.NoError(t, err)
	})
}
