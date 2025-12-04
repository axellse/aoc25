package main

//maybe inefficient?
import (
	"fmt"
	"os"
	"strings"
)

func compute(rows []string) (total int, removed [][2]int) {
	removed = [][2]int{}
	for ri, row := range rows {
		for ii, item := range row {
			if item == '.' {continue}

			adjacent := 0
			for _, cri := range []int{ri - 1, ri, ri + 1} {
				if cri == -1 || cri >= len(rows) {continue}
				for _, cii := range []int{ii - 1, ii, ii + 1} {
					if cii == -1 || cii >= len(row) {continue}
					
					if rows[cri][cii] == '@' && !(cii == ii && cri == ri) {
						adjacent++
					}
				}
			}

			if adjacent < 4 {
				total++
				removed = append(removed, [2]int{ri, ii})
			}
		}
	}
	return
}

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("input where?")
	}

	inputStr := strings.ReplaceAll(string(input), "\r", "")
	rows := strings.Split(inputStr, "\n")
	total := 0
	partOne := false
	for {
		tinc, dif := compute(rows)
		total += tinc

		if tinc == 0 || partOne {break}

		for _, d := range dif {
			ri, ii := d[0], d[1]
			rows[ri] = rows[ri][:ii] + "." + rows[ri][ii+1:]
		}
	}

	fmt.Println("total is", total)
}