package main

//pretty shit here two but runs quickly

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main2() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("input where?")
	}

	set := []string{}
	ranges := strings.Split(string(input), ",")
	for _, r := range ranges {
		nums := strings.Split(r, "-")
		n1, _ := strconv.Atoi(nums[0])
		n2, _ := strconv.Atoi(nums[1])
		for i := n1; i <= n2; i++ {
			set = append(set, strconv.Itoa(i))
		}
	}

	addedUp := 0
	alreadyFound := []string{} //need to keep track of identicals
	for _, num := range set {
		for i := 2; i <= len(num); i++ { //i är hur många bitar
			if len(num) % i != 0 {
				continue
			}
			pl := len(num) / i

			others := ""
			for v := 0; v <= len(num); v += pl {
				if len(num) < v + pl {
					continue
				}
				sec := num[v:v + pl]
				if others == "" {
					others = sec
					continue
				}

				if others == sec {
					continue
				} else {
					others = "INVALID"
					break
				}
			}

			if others == "INVALID" {
				continue
			} else {
				if !slices.Contains(alreadyFound, num) {
					//fmt.Println("found", num, "with length", pl, "section is", others)
					finale, _ := strconv.Atoi(num)
					addedUp += finale
					alreadyFound = append(alreadyFound, num)	
				}
			}
		}
	}
	fmt.Println("added up is", addedUp)
}