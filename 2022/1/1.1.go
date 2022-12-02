package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	fileContent, err := ioutil.ReadFile("puzzleInput.txt")

	if err != nil {
		fmt.Println(err)
	}

	allItems := strings.Split(string(fileContent), "\r\n\r\n")

	lastTotal := 0
	currentMax := 0
	for _, e := range allItems {
		currentTotal := 0
		backpack := strings.Split(e, "\r\n")
		for _, f := range backpack {
			n, _ := strconv.Atoi(f)
			currentTotal = currentTotal + n
		}

		if currentTotal > lastTotal {
			currentMax = currentTotal
			lastTotal = currentTotal
		}
	}

	fmt.Println(currentMax)
}
