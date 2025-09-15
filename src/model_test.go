package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBinomials(t *testing.T) {
	tests := []struct {
		value    string
		expected []string
	}{
		{"", []string{}},                       // empty string
		{"12", []string{"12"}},                 // 22
		{"1234", []string{"12", "34"}},         // 1234
		{"123456", []string{"12", "34", "56"}}, // 123456
		{"123", []string{"12", "3"}},           // 123
	}

	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			assert.Equal(t, tt.expected, GetBinomials(tt.value))
		})
	}
}

func TestIsBinomialDescribing(t *testing.T) {
	tests := []struct {
		value    string
		count    int
		figure   string
		expected bool
	}{
		{"14233221", 1, "4", true}, // one 4 in 14233221
		{"14233221", 2, "3", true}, // two 3s in 14233221
		{"14233221", 3, "2", true}, // three 2s in 14233221
		{"14233221", 2, "1", true}, // two 1s in 14233221
		{"22", 2, "2", true},       // two 2s in 22
		{"123456", 1, "2", true},   // one 2s in 123456
		{"123456", 3, "4", false},  // three 4s in 123456
	}

	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			assert.Equal(t, tt.expected, IsBinomialDescribing(tt.value, tt.count, tt.figure))
		})
	}
}

func TestIsEnoughBinomials(t *testing.T) {
	tests := []struct {
		value    string
		expected bool
	}{
		{"14233221", true}, // 14233221
		{"23322114", true}, // same as 14233221, but different order
		{"22", true},       // 22
		{"23", false},      // 23, there is 1 3s, not 2
		{"11", true},       // 11: there is 2 1s, not 1
		{"123", false},     // 123: not even length
		{"123456", false},  // 123456: not self-describing
		{"183110", false},  // 311018: not self-describing, we don't count 3s
		{"666666", false},  // 666666: only 6s, should be only one binomial
		{"", true},         // empty string
	}

	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			binomials := GetBinomials(tt.value)
			assert.Equal(t, tt.expected, IsEnoughBinomials(tt.value, binomials))
		})
	}
}

func TestAreBinomialsOrdered(t *testing.T) {
	tests := []struct {
		value    string
		expected bool
	}{
		{"14233221", true},  // 14233221
		{"23322114", false}, // same as 14233221, but different order
		{"22", true},        // 22
		{"666666", true},    // 666666: not self-describing
		{"", true},          // empty string
	}

	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			binomials := GetBinomials(tt.value)
			assert.Equal(t, tt.expected, AreBinomialsOrdered(binomials))
		})
	}
}
