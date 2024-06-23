package options

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLibraryOptions_SetLibraryNameAndVersion(t *testing.T) {
	t.Run("Should_Return_Error_When_Args_Are_Empty", func(t *testing.T) {
		// Given
		io := &LibraryOptions{}

		// When
		result := io.Set([]string{})

		// Then
		assert.Error(t, result)
	})
	t.Run("Should_Return_Error_When_Args_Aren't_Include_@", func(t *testing.T) {
		// Given
		io := &LibraryOptions{}

		// When
		result := io.Set([]string{"mockery"})

		// Then
		assert.Error(t, result)
	})
	t.Run("Success", func(t *testing.T) {
		// Given
		io := &LibraryOptions{}

		// When
		result := io.Set([]string{"github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.0"})

		// Then
		assert.Nil(t, result)
		assert.Equal(t, "github.com/golangci/golangci-lint/cmd/golangci-lint", io.Address)
		assert.Equal(t, "v1.59.0", io.Version)
		assert.Equal(t, "golangci-lint", io.LibName)
		assert.Equal(t, "github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.0", io.Package)
	})
}
