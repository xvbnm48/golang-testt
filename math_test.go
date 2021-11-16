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

func TestAdd_WithTestTable(t *testing.T) {
	testCases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{
			name:     "negative and negatice",
			a:        -1,
			b:        -1,
			expected: -2,
		}, {
			name:     "positive and negatice",
			a:        1,
			b:        -1,
			expected: 0,
		}, {
			name:     "positive and positive",
			a:        1,
			b:        1,
			expected: 2,
		}, {
			name:     "negative and positive",
			a:        -1,
			b:        1,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := Add(tc.a, tc.b)
			assert.Equal(t, tc.expected, c)
		})

	}
}

func TestAdd(t *testing.T) {
	t.Run("negative add negative", func(t *testing.T) {
		c := Add(-1, -2)
		assert.Equal(t, -3, c)
	})

	t.Run("positive add negative", func(t *testing.T) {
		c := Add(1, -2)
		assert.Equal(t, -1, c)
	})

}
