package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	inc := count_increases("puzzleInput.txt")
	fmt.Println(inc)
}

func count_increases(filePath string) int {
	fileContent, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Println(err)
	}

	depthMeasurments := strings.Split(string(fileContent), "\n")

	var lastDepth string
	var increases int

	for idx, depth := range depthMeasurments {

		if idx == 0 {
			lastDepth = depth
			continue
		}

		d, _ := strconv.Atoi(depth)
		ld, _ := strconv.Atoi(lastDepth)

		if d > ld {
			increases++
			lastDepth = depth
			continue
		}

		lastDepth = depth
	}

	return increases
}
