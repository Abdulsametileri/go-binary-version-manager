package installer

import (
	"testing"

	"github.com/Abdulsametileri/go-binary-version-manager/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestGetInstaller(t *testing.T) {
	t.Run("Failure", func(t *testing.T) {
		// Given

		// When
		installer, err := Get("unknownlib")

		// Then
		assert.Nil(t, installer)
		assert.Error(t, err)
	})
	t.Run("Success_For_Golangci_Lint", func(t *testing.T) {
		// Given

		// When
		installer, err := Get(model.LibraryGolangciLint.String())

		// Then
		assert.NotNil(t, installer)
		assert.Nil(t, err)
	})
	t.Run("Success_For_Mockery_Lint", func(t *testing.T) {
		// Given

		// When
		installer, err := Get(model.LibraryMockery.String())

		// Then
		assert.NotNil(t, installer)
		assert.Nil(t, err)
	})
}
