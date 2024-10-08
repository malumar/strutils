package strutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitQuotedString(t *testing.T) {
	assert.Equal(t, []string{`123`, "456", "jureczek abc", "hello world", `item"`}, SplitQuotedString(`123 456 "jureczek abc" "hello world"item"`))
	assert.Equal(t, []string{"hello world", `item"accd`}, SplitQuotedString(`"hello world"item"accd`))
	assert.Equal(t, []string{"hello world", `item\"accd`}, SplitQuotedString(`"hello world"item\"accd`))
}
