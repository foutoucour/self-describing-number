package main

import (
	"fmt"
	//"fmt"
	"sort"
	"strconv"
	"sync"
	"time"

	//"time"
	"github.com/alecthomas/kong"
)

// isValid checks if a number is self-describing.
func isValid(number string) bool {
	if _, err := strconv.Atoi(number); err != nil {
		return false
	}

	if len(number)%2 != 0 {
		return false
	}

	binomials := GetBinomials(number)
	if len(binomials) != len(RemoveDuplicates(binomials)) {
		return false
	}
	if !IsEnoughBinomials(number, binomials) {
		return false
	}
	if !AreBinomialsOrdered(binomials) {
		return false
	}

	for _, binomial := range binomials {
		count, _ := strconv.Atoi(string(binomial[0]))
		figure := string(binomial[1])
		if !IsBinomialDescribing(number, count, figure) {
			return false
		}
	}

	return true
}

// processTask processes a list of numbers and returns the valid ones.
func processTask(task []int, wg *sync.WaitGroup, results *[]string, mu *sync.Mutex) {
	defer wg.Done()
	var validNumbers []string
	for _, i := range task {
		if isValid(strconv.Itoa(i)) {
			validNumbers = append(validNumbers, strconv.Itoa(i))
		}
	}
	mu.Lock()
	*results = append(*results, validNumbers...)
	mu.Unlock()
}

// command orchestrates the processing of numbers using multiple goroutines.
func command(number int) []string {
	numbers := make([]int, number)
	for i := 0; i < number; i++ {
		numbers[i] = i
	}

	tasks := RoundRobinSublists(numbers, 20)
	var results []string
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, task := range tasks {
		wg.Add(1)
		go processTask(task, &wg, &results, &mu)
	}

	wg.Wait()
	sort.Slice(results, func(i, j int) bool {
		ni, _ := strconv.Atoi(results[i])
		nj, _ := strconv.Atoi(results[j])
		return ni < nj
	})
	return results
}

func run(number int) {
	//number := 100000000 // Example input

	start := time.Now()
	results := command(number)
	end := time.Now()

	fmt.Printf("Self-describing numbers up to %d:\n", number)
	for _, result := range results {
		fmt.Println(result)
	}
	fmt.Printf("Found %d self-describing numbers in %.2f seconds\n", len(results), end.Sub(start).Seconds())
}

var CLI struct {
	Run struct {
		Number int `arg:"" name:"number" help:"number to run against" type:"int"`
	} `cmd:"" help:"Find the self-describing numbers."`
}

func main() {
	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "run <number>":
		run(CLI.Run.Number)
	default:
		panic(ctx.Command())
	}
}
