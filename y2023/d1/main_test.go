package d1

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	f, err := os.Open("test-input-1.txt")
	assert.NoError(t, err)

	want := 142
	got, err := partOne(f)
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestPartTwo(t *testing.T) {
	f, err := os.Open("test-input-2.txt")
	assert.NoError(t, err)

	want := 281
	got, err := partTwo(f)
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
