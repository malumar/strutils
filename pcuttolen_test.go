package strutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCutStringToLen(t *testing.T) {
	fullVal := "john doe"
	assert.Equal(t, "john", CutStringToLen(fullVal, 4))
	assert.Equal(t, fullVal, CutStringToLen(fullVal, 50))
	assert.Equal(t, fullVal, CutStringToLen(fullVal, 0))
	assert.Equal(t, fullVal, CutStringToLen(fullVal, -1))
}
