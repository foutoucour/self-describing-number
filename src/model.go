package main

import (
	"sort"
	"strings"
)

// IsLengthEven checks if the length of the number is even.
func IsLengthEven(number string) bool {
	return len(number)%2 == 0
}

func RemoveDuplicates(input []string) []string {
	seen := make(map[string]struct{})
	var result []string

	for _, value := range input {
		if _, exists := seen[value]; !exists {
			seen[value] = struct{}{}
			result = append(result, value)
		}
	}

	return result
}

// GetBinomials splits the number into pairs of two characters.
func GetBinomials(number string) []string {
	binomials := []string{}
	for i := 0; i < len(number); i += 2 {
		end := i + 2
		if end > len(number) {
			end = len(number)
		}
		binomials = append(binomials, number[i:end])
	}
	return binomials
}

// IsBinomialDescribing checks if the figure appears count times in the number.
func IsBinomialDescribing(number string, count int, figure string) bool {
	return strings.Count(number, figure) == count
}

// IsEnoughBinomials checks if the number of unique figures matches the number of binomials.
func IsEnoughBinomials(number string, binomials []string) bool {
	uniqueFigures := map[rune]struct{}{}
	for _, c := range number {
		uniqueFigures[c] = struct{}{}
	}
	return len(uniqueFigures) == len(binomials)
}

// AreBinomialsOrdered checks if the figures in the binomials are in descending order.
func AreBinomialsOrdered(binomials []string) bool {
	figures := []string{}
	for _, b := range binomials {
		if len(b) > 1 {
			figures = append(figures, string(b[1]))
		}
	}
	sortedFigures := append([]string{}, figures...)
	sort.Sort(sort.Reverse(sort.StringSlice(sortedFigures)))
	return equalSlices(figures, sortedFigures)
}

// equalSlices checks if two slices are equal.
func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// roundRobinSublists splits a range of numbers into sublists for round-robin distribution.
func RoundRobinSublists(numbers []int, n int) [][]int {
	sublists := make([][]int, n)
	for i, num := range numbers {
		sublists[i%n] = append(sublists[i%n], num)
	}
	return sublists
}
