package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func partOne(lines []string) {
	firstCol := []int64{}
	secondCol := []int64{}

	for _, line := range lines {
		words := strings.Split(line, "   ")
		if len(words) < 2 {
			continue
		}
		first, _ := strconv.ParseInt(words[0], 10, 64)
		second, _ := strconv.ParseInt(words[1], 10, 64)

		firstCol = append(firstCol, first)
		secondCol = append(secondCol, second)
	}
	sort.Slice(firstCol, func(i, j int) bool {
		return firstCol[i] < firstCol[j]
	})
	sort.Slice(secondCol, func(i, j int) bool {
		return secondCol[i] < secondCol[j]
	})
	distance := int64(0)

	for i := 0; i < len(firstCol); i++ {
		localDistance := float64(firstCol[i] - secondCol[i])
		distance += int64(math.Abs(localDistance))
	}
	fmt.Println("Distance: ")
	fmt.Println(distance)
}

func partTwo(lines []string) {

	firstCol := []int64{}
	secondColMap := map[int64]int64{}

	for _, line := range lines {
		words := strings.Split(line, "   ")
		if len(words) < 2 {
			continue
		}
		first, _ := strconv.ParseInt(words[0], 10, 64)
		second, _ := strconv.ParseInt(words[1], 10, 64)
		firstCol = append(firstCol, first)
		if _, ok := secondColMap[second]; ok {
			secondColMap[second]++
		} else {
			secondColMap[second] = 1
		}
	}

	similarity := int64(0)
	for i := 0; i < len(firstCol); i++ {
		if _, ok := secondColMap[firstCol[i]]; ok {
			similarity += secondColMap[firstCol[i]] * firstCol[i]
		}
	}
	fmt.Println("Similarity: ")
	fmt.Println(similarity)
}

// main function
func main() {
	file, _ := os.ReadFile("./2024/01/input.txt")
	lines := strings.Split(string(file), "\n")
	partOne(lines)
	partTwo(lines)

}
