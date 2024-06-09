package options

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLibraryOptions_SetLibraryNameAndVersion(t *testing.T) {
	t.Run("Should_Return_Error_When_Args_Are_Empty", func(t *testing.T) {
		// Given
		io := &LibraryOptions{}

		// When
		result := io.SetLibraryNameAndVersion([]string{})

		// Then
		assert.Error(t, result)
	})
	t.Run("Should_Return_Error_When_Args_Aren't_Include_@", func(t *testing.T) {
		// Given
		io := &LibraryOptions{}

		// When
		result := io.SetLibraryNameAndVersion([]string{"mockery"})

		// Then
		assert.Error(t, result)
	})
	t.Run("Success", func(t *testing.T) {
		// Given
		io := &LibraryOptions{}

		// When
		result := io.SetLibraryNameAndVersion([]string{"mockery@v2.20.0"})

		// Then
		assert.Nil(t, result)
		assert.Equal(t, "mockery", io.LibraryName)
		assert.Equal(t, "v2.20.0", io.Version)
	})
}
