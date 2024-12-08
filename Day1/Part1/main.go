package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseInput(filename string) ([]int, []int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	var leftList, rightList []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Invalid input line:", line)
			continue
		}

		left, err1 := strconv.Atoi(parts[0])
		right, err2 := strconv.Atoi(parts[1])

		if err1 != nil || err2 != nil {
			fmt.Println("Error parsing numbers:", err1, err2)
			continue
		}

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	return leftList, rightList
}

func calculateTotalDistance(left []int, right []int) int {
	totalDistance := 0

	for i := 0; i < len(left); i++ {
		distance := int(math.Abs(float64(left[i] - right[i])))
		totalDistance += distance
	}

	return totalDistance
}

func main() {
	leftList, rightList := parseInput(os.Args[1])

	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDistance := calculateTotalDistance(leftList, rightList)
	fmt.Println("Total Distance:", totalDistance)
}
