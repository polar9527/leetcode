package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	s := "ac"
	result := longestPalindrome(s)
	assert.Equal(t, "c", result)
}
