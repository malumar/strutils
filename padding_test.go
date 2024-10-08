package strutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLeftPad(t *testing.T) {
	assert.Equal(t, "0001", LeftPad2Len("1", "0", 4))
	assert.Equal(t, "0011", LeftPad2Len("11", "0", 4))
}
