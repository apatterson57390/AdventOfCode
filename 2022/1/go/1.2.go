package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {

	fileContent, err := ioutil.ReadFile("puzzleInput.txt")

	if err != nil {
		fmt.Println(err)
	}

	allItems := strings.Split(string(fileContent), "\r\n\r\n")

	allTotals := []int{}
	for _, e := range allItems {
		currentTotal := 0
		backpack := strings.Split(e, "\r\n")
		for _, f := range backpack {
			n, _ := strconv.Atoi(f)
			currentTotal = currentTotal + n
		}

		allTotals = append(allTotals, currentTotal)
	}

	sort.Ints(allTotals)

	var result int
	for _, t := range allTotals[len(allTotals)-3:] {
		result = result + t
	}

	fmt.Println(result)
}
