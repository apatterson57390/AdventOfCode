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

	var (
		slidingWindow []int
		lastWindowSum int = 0
		increases     int = 0
	)

	for _, depth := range depthMeasurments {

		// Continue to add to the sliding window until we reach 3 numbers
		if len(slidingWindow) < 3 {
			d, _ := strconv.Atoi(depth)
			slidingWindow = append(slidingWindow, d)
		}

		// reset/initialize the current sliding window sum before adding
		windowSum := 0

		if len(slidingWindow) == 3 {
			// Add together the 3 numbers
			for _, dv := range slidingWindow {
				windowSum += dv
			}

			// Account for the first iteration which won't have a lastWindowSum and which won't count towards increases
			if lastWindowSum == 0 {
				lastWindowSum = windowSum
				slidingWindow = slidingWindow[1:]
				continue
			} else {
				// Determine if there was a increase from the last go around
				if windowSum > lastWindowSum {
					increases++
				}

				lastWindowSum = windowSum
				// Pop the first element off of the slidingWindow to "slide"
				slidingWindow = slidingWindow[1:]
			}
		}
	}

	return increases
}
