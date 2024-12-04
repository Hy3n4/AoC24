package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func readLines(path string) ([]string, error) {
	inputFile, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening input file: %v", err)
	}
	defer inputFile.Close()

	var lines []string

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func deleteElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	var column1 []int
	var column2 []int
	var finalSlice []float64

	for _, line := range lines {
		r, _ := regexp.Compile(`^(\d{5})\s+(\d{5})$`)
		match := r.FindStringSubmatch(line)
		int1, err := strconv.Atoi(match[1])
		if err != nil {
			fmt.Printf("error converting first column nummber %v", err)
		}
		int2, err := strconv.Atoi(match[2])
		if err != nil {
			fmt.Printf("error converting second column nummber %v", err)
		}
		column1 = append(column1, int1)
		column2 = append(column2, int2)

		sort.Ints(column1)
		sort.Ints(column2)

		finalSlice += math.Abs(float64(column1[0]) - float64(column2[0]))
		column1 = deleteElement(column1, 0)
		column2 = deleteElement(column2, 0)
	}
	fmt.Println(finalSlice)
}
