package enabler

import (
	"github.com/Abdulsametileri/go-binary-version-manager/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGet(t *testing.T) {
	t.Run("Failure", func(t *testing.T) {
		// Given

		// When
		enabler, err := Get("unknownlib")

		// Then
		assert.Nil(t, enabler)
		assert.Error(t, err)
	})
	t.Run("Success_For_Golangci_Lint", func(t *testing.T) {
		// Given

		// When
		enabler, err := Get(model.LibraryGolangciLint.String())

		// Then
		assert.NotNil(t, enabler)
		assert.Nil(t, err)
	})
}
