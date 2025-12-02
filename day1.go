package main

//pretty ass here
import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func two(lines []string) int {
	pos := 50
	zeroesPassed := 0

	for _, line := range lines {
		distance, _ := strconv.Atoi(line[1:])
		d := line[0]
		for range distance {
			if d == 'L' {
				pos--
			} else {
				pos++
			}

			if pos%100 == 0 {
				zeroesPassed++
			}
		}
	}
	return zeroesPassed
}

func main1() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("input where?")
	}

	inputStr := strings.ReplaceAll(string(input), "\r", "")
	pos := 50
	zeroes := 0
	lines := strings.Split(inputStr, "\n")

	for _, line := range lines {
		distance, _ := strconv.Atoi(line[1:])
		result := pos + distance
		if line[0] == 'L' {
			result = pos - distance	
		}

		for {
			if result < 0 {
				result = 100 + result
			} else if result > 99 {
				result = 0 + (result - 100)
			} else {
				break
			}
		}

		if result == 0 {
			zeroes++
		}
		pos = result
	}
	
	fmt.Println("ended up at", pos, "moved past zero", two(lines), "times", "landed on zero", zeroes, "times")
}