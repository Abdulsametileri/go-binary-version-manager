package installer

import (
	"context"
	"errors"
	"github.com/Abdulsametileri/go-binary-version-manager/internal/mocks"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_mockeryInstaller_Install(t *testing.T) {
	t.Run("Should_Return_Error_When_Go_GoPath_Command_Failed", func(t *testing.T) {
		// Given
		expectedErr := errors.New("command err")
		runner := mocks.NewCommandRunner(t)
		runner.On("RunWith", "env", "GOPATH").Return("", expectedErr)
		sut := &mockeryInstaller{goCmdRunner: runner}

		// When
		result := sut.Install(context.Background(), "v2.20.0")

		// Then
		assert.ErrorIs(t, result, expectedErr)
		runner.AssertExpectations(t)
	})
	t.Run("Should_Return_Error_When_Go_GoInstall_Command_Failed", func(t *testing.T) {
		// Given
		expectedErr := errors.New("command err")
		runner := mocks.NewCommandRunner(t)
		runner.On("RunWith", "env", "GOPATH").Return("/Users/samet.ileri/go", nil)
		runner.On("RunWith", "install", "github.com/vektra/mockery/v2@v2.20.0").Return("", expectedErr)

		sut := &mockeryInstaller{goCmdRunner: runner, goBinEnvKey: "GOBIN_TEST"}

		// When
		result := sut.Install(context.Background(), "v2.20.0")

		// Then
		assert.ErrorIs(t, result, expectedErr)
		runner.AssertExpectations(t)
	})
	t.Run("Success", func(t *testing.T) {
		// Given
		runner := mocks.NewCommandRunner(t)
		runner.On("RunWith", "env", "GOPATH").Return("/Users/samet.ileri/go", nil)
		runner.On("RunWith", "install", "github.com/vektra/mockery/v2@v2.20.0").Return("INSTALLED", nil)
		os.Setenv("GOBIN_TEST", "GOBINTESTVALUE_BEFORE_TEST")

		sut := &mockeryInstaller{goCmdRunner: runner, goBinEnvKey: "GOBIN_TEST"}

		// When
		result := sut.Install(context.Background(), "v2.20.0")

		// Then
		assert.NoError(t, result)
		assert.Equal(t, "GOBINTESTVALUE_BEFORE_TEST", os.Getenv("GOBIN_TEST"))
		runner.AssertExpectations(t)
	})
}
