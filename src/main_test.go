package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		value    string
		expected bool
	}{
		{"14233221", true},
		{"23322114", false}, // same as 14233221, but different order
		{"22", true},
		{"23", false},     // there is 1 3s, not 2
		{"11", false},     // there is 2 1s, not 1
		{"123", false},    // not even length
		{"123456", false}, // not self-describing
		{"311018", false}, // not self-describing, we don't count 1s
		{"666666", false}, // not self-describing
		{"", false},       // empty string
	}

	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			assert.Equal(t, tt.expected, isValid(tt.value))
		})
	}
}

func TestCommand(t *testing.T) {
	tests := []struct {
		value    int
		expected []string
	}{
		{0, []string(nil)},         // 0ms
		{10, []string(nil)},        // 0ms
		{100, []string{"22"}},      // 0ms
		{1000, []string{"22"}},     // 0ms
		{10000, []string{"22"}},    // 8ms
		{100000, []string{"22"}},   // 23ms
		{1000000, []string{"22"}},  // 1s
		{10000000, []string{"22"}}, // 3s
		{100000000, []string{
			"22", "14233221", "14331231", "14333110", "15143331", "15233221",
			"15331231", "15333110", "16143331", "16153331", "16233221",
			"16331231", "16333110", "17143331", "17153331", "17163331",
			"17233221", "17331231", "17333110", "18143331", "18153331",
			"18163331", "18173331", "18233221", "18331231", "18333110",
			"19143331", "19153331", "19163331", "19173331", "19183331",
			"19233221", "19331231", "19333110", "23322110", "33123110",
		}}, // 2mins
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.value), func(t *testing.T) {
			assert.Equal(t, tt.expected, command(tt.value))
		})
	}
}
