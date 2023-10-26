package math_test

import (
	"testing"
	"time"
	. "unit-test-golang-v2"

	"github.com/stretchr/testify/assert"
)

func TestAdd_WithTestTable(t *testing.T) {
	testCases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{
			name:     "negative and negative",
			a:        -1,
			b:        -1,
			expected: -2,
		},
		{

			name:     "negative and positive",
			a:        -1,
			b:        1,
			expected: 0,
		},
		{

			name:     "positive and negative",
			a:        1,
			b:        1,
			expected: 2,
		},
		{

			name:     "positive and positive",
			a:        1,
			b:        -1,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := Add(tc.a, tc.b)
			assert.Equal(t, tc.expected, e)
		})
	}
}

func TestAbsolute(t *testing.T) {
	t.Skip()
	t.Run("negative test case", func(t *testing.T) {
		c := Absolute(-5)
		assert.Equal(t, 5, c)
	})
	t.Run("positive test case", func(t *testing.T) {
		c := Absolute(5)
		assert.Equal(t, 5, c)
	})
}

func TestAdd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	<-time.After(5 * time.Second)

	t.Run("negative add negative", func(t *testing.T) {
		c := Add(-1, -2)
		assert.Equal(t, -3, c)
	})
	t.Run("positive add negative", func(t *testing.T) {
		c := Add(1, -2)
		assert.Equal(t, -1, c)
	})
}
