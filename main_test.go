package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestCompareStrings(t *testing.T) {
	s1 := `  123
	456
	789  `
	s2 := `AAA
	456
	`
	exp := []string{
		"AAA",
	}
	got := compareStrings(s1, s2)
	assert.Equal(t, exp, got)

	exp = []string{
		"123",
		"789",
	}
	got = compareStrings(s2, s1)
	assert.Equal(t, exp, got)
}
