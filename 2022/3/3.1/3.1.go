package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func checkItems(items []string) int {

	for i := 0; i < len(items); i++ {

		// split item into two halfs
		halfCount := len(items[i]) / 2
		firstHalf := strings.Split(items[i], "")[0:halfCount]
		secondHalf := strings.Split(items[i], "")[halfCount:len(items[i])]

		item := map[string][]string{"firstHalf": firstHalf, "secondHalf": secondHalf}

		// match letters
		letterMap := map[string]string{}
		for f, s := range strings.Split(items[i], "") {
			lettermap
		}

	}
	return 0
}

func main() {
	input, err := ioutil.ReadFile("puzzleInput.txt")

	if err != nil {
		log.Fatal(err)
	}

	checkItems(strings.Split(string(input), "\n"))
}
