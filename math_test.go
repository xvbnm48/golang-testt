package go_testing_baru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMath(t *testing.T) {
	t.Run("negatie test", func(t *testing.T) {
		c := Math(-5)
		assert.Equal(t, 5, c)

	})
	t.Run("testing postivie", func(t *testing.T) {
		c := Math(5)
		assert.Equal(t, 5, c)

	})
}
