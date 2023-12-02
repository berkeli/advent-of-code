package d1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	want := 142
	got, err := partOne("test-input-1.txt")
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestPartTwo(t *testing.T) {
	want := 281
	got, err := partTwo("test-input-2.txt")
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
