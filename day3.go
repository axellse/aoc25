package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func IndexOf(haystack []int, needle int) int {
    for i, v := range haystack {
        if v == needle {
            return i
        }
    }
    return -1
}

func main3() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("input where?")
	}

	inputStr := strings.ReplaceAll(string(input), "\r", "")
	banks := strings.Split(string(inputStr), "\n")
	total := 0
	for _, b := range banks {
		bank := []int{}
		for _, v := range b {
			i, _ := strconv.Atoi(string(v))
			bank = append(bank, i)
		}

		spaceRequired := 12 //use 2 for part 1 and 12 for part 2
		startDigit := 0

		final := ""
		for {
			if spaceRequired == 0 {break}
			rang := bank[startDigit:len(bank)-spaceRequired +1]
			val := slices.Max(rang)
			startDigit += IndexOf(rang, val) + 1

			fmt.Println("picked", val, "new from", startDigit, "to", len(bank)-spaceRequired +1, "aka", rang)
			final += strconv.Itoa(val)
			spaceRequired--
		}
		fmt.Println("final is", final)
		finalAddedUp, _ := strconv.Atoi(final)
		total += finalAddedUp
	}
	fmt.Println("final result for all is", total)
}