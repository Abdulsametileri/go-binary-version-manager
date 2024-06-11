package installer

import (
	"context"
	"errors"
	"testing"

	"github.com/Abdulsametileri/go-binary-version-manager/internal/installer/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_golangciLintInstaller_Install(t *testing.T) {
	t.Run("Failure", func(t *testing.T) {
		// Given
		expectedErr := errors.New("runner err")
		runner := mocks.NewCommandRunner(t)
		runner.On("RunWith", mock.Anything).Return("", expectedErr)
		sut := golangciLintInstaller{curlCmdRunner: runner}

		// When
		err := sut.Install(context.Background(), "golangci-lint@v1.55.0")

		// Then
		assert.ErrorIs(t, err, expectedErr)
		runner.AssertExpectations(t)
	})
	t.Run("Success", func(t *testing.T) {
		// Given
		runner := mocks.NewCommandRunner(t)

		expectedCmd := "curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin/glvm/golangci-lint/golangci-lint@v1.55.0 golangci-lint@v1.55.0"

		runner.On("RunWith", expectedCmd).Return("out", nil)
		sut := golangciLintInstaller{curlCmdRunner: runner}

		// When
		err := sut.Install(context.Background(), "golangci-lint@v1.55.0")

		// Then
		assert.Nil(t, err)
		runner.AssertExpectations(t)
	})
}
