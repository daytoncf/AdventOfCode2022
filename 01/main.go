package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part one result: \n%v\n\n", partOne())
	fmt.Printf("Part two result: \n%v\n", partTwo())
}

// convert file into string
func fileToString(filename string) string {
	file, err := os.ReadFile(filename)
	checkErr(err)
	return string(file)
}

// for repeated error checking functionality
func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func partOne() string {
	var contents string = fileToString("./input.txt")

	arr := strings.Split(contents, "\n")

	// Utilizing a map w/ int array as value in this approach to hold onto as many values as possible for part 2
	elfMap := make(map[int][]int)
	var elfNum int = 1

	// Iterate across the list to build out the map
	for _, v := range arr {
		if v != "" {
			itemCals, err := strconv.Atoi(v)
			checkErr(err)
			elfMap[elfNum] = append(elfMap[elfNum], itemCals)
		} else {
			elfNum++
		}
	}

	var maxCals int = 0
	var maxElf int = 0

	// Iterate over the map to read each elfs list of food items
	for k, v := range elfMap {
		var sum int = 0
		// Iterate over each list of food items to find max
		for _, num := range v {
			sum += num
		}
		if sum > maxCals {
			maxCals = sum
			maxElf = k
		}
	}

	return fmt.Sprintf("Elf carrying the most is Elf %v.\nThey are carrying %v calories worth of food.\n", maxElf, maxCals)
}

func partTwo() string {
	var contents string = fileToString("./input.txt")

	arr := strings.Split(contents, "\n")

	// Initialize max most calories
	var m1, m2, m3 int = 0, 0, 0

	// Sum for the current elf's total calories carried
	var elfsTotalCals int = 0

	// Iterate across list of values
	for _, v := range arr {
		if v != "" {
			itemCals, err := strconv.Atoi(v)
			checkErr(err)
			elfsTotalCals += itemCals
		} else {
			// Compare elf's total calories to the top 3
			if elfsTotalCals > m1 {
				m1 = elfsTotalCals
			} else if elfsTotalCals > m2 {
				m2 = elfsTotalCals
			} else if elfsTotalCals > m3 {
				m3 = elfsTotalCals
			}

			// Moving onto next elf, reset
			elfsTotalCals = 0
		}
	}
	return fmt.Sprintf("Top 3 elves are carrying %v, %v, %v in descending order.\nTotal is %v\n", m1, m2, m3, (m1 + m2 + m3))
}
