package lister

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGet(t *testing.T) {
	t.Run("Failure", func(t *testing.T) {
		// Given

		// When
		lister, err := Get("unknown")

		// Then
		assert.Nil(t, lister)
		assert.Error(t, err)
	})
	t.Run("Success_For_Stdout", func(t *testing.T) {
		// Given

		// When
		lister, err := Get("stdout")

		// Then
		assert.NotNil(t, lister)
		assert.Nil(t, err)
	})
}
