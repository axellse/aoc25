package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("input where?")
	}

	inputStr := strings.ReplaceAll(string(input), "\r", "")
	sections := strings.Split(inputStr, "\n\n")
	ranges := [][2]int{} //first number, second number

	for rang := range strings.SplitSeq(sections[0], "\n") {
		nums := strings.Split(rang, "-")
		n1, _ := strconv.Atoi(nums[0])
		n2, _ := strconv.Atoi(nums[1])
		ranges = append(ranges, [2]int{n1, n2})
	}

	slices.SortFunc(ranges, func(a [2]int, b [2]int) int {
		return a[0] - b[0]
	})
	result := [][2]int{ranges[0]}
	for _, rang := range ranges[1:] {
		prevRang := result[len(result) - 1]
		n1 := rang[0]
		n2 := rang[1]
		if n1 <= prevRang[1] {
			end := max(prevRang[1], n2)
			result[len(result) -1][1] = end
		} else {
			result = append(result, [2]int{n1, n2})
		}
	}

	allFreshIds := 0
	for _, v := range result {
		allFreshIds += v[1] - v[0] + 1
	}

	freshIds := 0
	for idStr := range strings.SplitSeq(sections[1], "\n") {
		id, _ := strconv.Atoi(idStr)
		for _, rang := range ranges {
			if id >= rang[0] && id <= rang[1] {
				freshIds++
				break
			}
		}
	}

	fmt.Println("total available fresh ids (part 1)", freshIds)
	fmt.Println("total fresh ids (part 2)", allFreshIds)
}