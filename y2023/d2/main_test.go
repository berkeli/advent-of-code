package d2

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	f, err := os.Open("test-input.txt")
	assert.NoError(t, err)

	want := 8
	got, err := partOne(f)
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestPartTwo(t *testing.T) {
	f, err := os.Open("test-input.txt")
	assert.NoError(t, err)

	want := 2286
	got, err := partTwo(f)
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
