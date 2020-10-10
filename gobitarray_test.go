package gobitarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetGet(t *testing.T) {
	arr := New(5)

	arr.Set(2)
	arr.Set(4)

	var tests = []struct {
		position int
		expected int
	}{
		{0, 0},
		{1, 0},
		{2, 1},
		{3, 0},
		{4, 1},
	}

	for _, tt := range tests {
		actual, _ := arr.Get(tt.position)
		assert.Equal(t, tt.expected, actual)
	}
}

func TestUnset(t *testing.T) {
	arr := New(5)

	arr.Set(2)

	actual, _ := arr.Get(2)
	assert.Equal(t, 1, actual)

	arr.Unset(2)
	lastActual, _ := arr.Get(2)
	assert.Equal(t, 0, lastActual)
}

func TestToggle(t *testing.T) {
	arr := New(5)

	arr.Set(2)

	actual, _ := arr.Get(2)
	assert.Equal(t, 1, actual)

	arr.Toggle(2)
	actual2, _ := arr.Get(2)
	assert.Equal(t, 0, actual2)

	arr.Toggle(2)
	actual3, _ := arr.Get(2)
	assert.Equal(t, 1, actual3)
}

func TestReset(t *testing.T) {
	arr := New(5)

	arr.Set(2)
	arr.Set(3)

	actual2, _ := arr.Get(2)
	assert.Equal(t, 1, actual2)

	actual3, _ := arr.Get(3)
	assert.Equal(t, 1, actual3)

	arr.Reset()

	lastActual2, _ := arr.Get(2)
	assert.Equal(t, 0, lastActual2)

	lastActual3, _ := arr.Get(3)
	assert.Equal(t, 0, lastActual3)
}
